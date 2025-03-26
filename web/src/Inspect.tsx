import { Clock, Copy } from 'lucide-react'
import { useEffect, useState } from 'react'
import { useParams } from 'react-router'
import './App.css'
import Footer from './components/Footer'
import Header from './components/Header'

interface Request {
  id: string
  method: string
  path: string
  headers: object
  timestamp: string
  form: object
  body: object
  data: string
}

function Inspect() {
  const { slug } = useParams()

  const slugUrl = `${location.protocol}://${location.host}/${slug}`

  const [showToast, setShowToast] = useState(false)

  const [requests, setRequests] = useState<Request[]>([])

  useEffect(() => {
    fetch(`/api/${slug}/inspect`)
      .then((res) => res.json())
      .then((data) => {
        setRequests(
          data.result.map((item) => {
            return {
              ...item.request,
              id: item._id,
              timestamp: item.created,
            }
          }),
        )
      })
      .catch((err) => {
        console.log(err.message)
      })
  }, [])

  useEffect(() => {
    const sse = new EventSource('/api/sse?stream=requests', {
      withCredentials: true,
    })

    function getRealtimeData(data) {
      setRequests((prev) => {
        return [
          {
            ...data.request,
            id: data._id,
            timestamp: data.created,
          },
          ...prev,
        ]
      })
    }
    sse.onmessage = (e) => getRealtimeData(JSON.parse(e.data))
    sse.onerror = (e) => {
      console.log(`sse error: ${JSON.stringify(e)}`)

      sse.close()
    }

    return () => {
      sse.close()
    }
  }, [])

  const copyUrl = () => {
    navigator.clipboard.writeText(slugUrl)
    setShowToast(true)
    setTimeout(() => setShowToast(false), 3000)
  }

  const getMethodColor = (method) => {
    const colors = {
      GET: 'bg-blue-100 text-blue-800',
      POST: 'bg-green-100 text-green-800',
      PUT: 'bg-yellow-100 text-yellow-800',
      DELETE: 'bg-red-100 text-red-800',
    }
    return colors[method] || 'bg-gray-100 text-gray-800'
  }

  return (
    <div className='min-h-screen bg-gray-50'>
      {/* Toast Notification */}
      <div
        className={`fixed top-4 right-4 transform transition-all duration-300 ${
          showToast
            ? 'translate-y-0 opacity-100'
            : '-translate-y-4 opacity-0 pointer-events-none'
        }`}
      >
        <div className='bg-gray-800 text-white px-4 py-2 rounded-lg shadow-lg flex items-center'>
          <svg
            className='w-5 h-5 mr-2 text-green-400'
            fill='none'
            strokeLinecap='round'
            strokeLinejoin='round'
            strokeWidth='2'
            viewBox='0 0 24 24'
            stroke='currentColor'
          >
            <path d='M5 13l4 4L19 7'></path>
          </svg>
          URL copied to clipboard!
        </div>
      </div>

      <Header />

      <main className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8'>
        {/* Bin URL */}
        <div className='bg-white rounded-lg shadow p-6 mb-8'>
          <div className='flex items-center justify-between'>
            <div className='flex-1'>
              <h2 className='text-lg font-medium text-gray-900'>
                Your Bin URL
              </h2>
              <div className='mt-2 flex items-center'>
                <code className='flex-1 block text-sm font-mono bg-gray-50 p-3 rounded-md'>
                  {slugUrl}
                </code>
                <button
                  onClick={copyUrl}
                  className='ml-3 inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500'
                >
                  <Copy className='h-4 w-4 mr-1' />
                  Copy
                </button>
              </div>
            </div>
          </div>
        </div>

        {/* Requests List and Details */}
        <div className='bg-white rounded-lg shadow'>
          <div className='p-4 border-b border-gray-200 flex items-center justify-between'>
            <h2 className='text-lg font-medium text-gray-900'>
              Recent Requests
            </h2>
          </div>
          <div className='divide-y divide-gray-200'>
            {requests.map((request) => (
              <div
                key={request.id}
                className='group bg-white transition-all duration-500'
              >
                <div className='p-4'>
                  <div className='flex items-center justify-between'>
                    <div className='flex items-center space-x-4'>
                      <span
                        className={`px-2.5 py-0.5 rounded-md text-sm font-medium ${getMethodColor(
                          request.method,
                        )}`}
                      >
                        {request.method}
                      </span>
                      <span className='text-sm font-mono text-gray-900'>
                        {request.path}
                      </span>
                    </div>
                    <span className='flex items-center text-sm text-gray-500 pr-3'>
                      <Clock className='h-4 w-4 mr-1' />
                      {new Date(request.timestamp).toLocaleString()}
                    </span>
                  </div>

                  {/* Request Details */}
                  <div className='mt-4 space-y-4'>
                    {/* Headers */}
                    <div>
                      <h3 className='text-sm font-medium text-gray-500 mb-2'>
                        Headers
                      </h3>
                      <div className='bg-white rounded-md overflow-hidden shadow-sm'>
                        <div className='divide-y divide-gray-200'>
                          {Object.entries(request.headers).map(
                            ([key, value], idx) => (
                              <div
                                key={idx}
                                className='flex px-3 py-2 hover:bg-gray-50'
                              >
                                <span className='w-1/3 font-mono text-sm text-gray-600'>
                                  {key}
                                </span>
                                <span className='w-2/3 font-mono text-sm text-balance text-gray-900 break-all'>
                                  {value}
                                </span>
                              </div>
                            ),
                          )}
                        </div>
                      </div>
                    </div>

                    {request.form && (
                      <div>
                        <h3 className='text-sm font-medium text-gray-500 mb-2'>
                          Form Body
                        </h3>
                        <div className='bg-gray-50 rounded-md p-4 overflow-x-auto font-mono text-sm'>
                          {JSON.stringify(request.form, null, 2)}
                        </div>
                      </div>
                    )}

                    {request.body && (
                      <div>
                        <h3 className='text-sm font-medium text-gray-500 mb-2'>
                          JSON Body
                        </h3>
                        <div className='bg-gray-50 rounded-md p-4 overflow-x-auto font-mono text-sm'>
                          {JSON.stringify(request.body, null, 2)}
                        </div>
                      </div>
                    )}

                    {!request.form && !request.body && request.data && (
                      <div>
                        <h3 className='text-sm font-medium text-gray-500 mb-2'>
                          Raw Body
                        </h3>
                        <div className='bg-gray-50 rounded-md p-4 overflow-x-auto font-mono text-sm'>
                          {request.data}
                        </div>
                      </div>
                    )}

                    {!request.data && (
                      <div>
                        <h3 className='text-sm font-medium text-gray-500 mb-2'>
                          No Body
                        </h3>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>
      </main>

      <Footer></Footer>
    </div>
  )
}

export default Inspect
