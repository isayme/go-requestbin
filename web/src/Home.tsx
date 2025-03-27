import { ArrowRight } from 'lucide-react'
import { customAlphabet } from 'nanoid'
import { useEffect } from 'react'
import { useNavigate } from 'react-router'
import Footer from './components/Footer'
import Header from './components/Header'
import { appName } from './global'

const nanoid = customAlphabet('1234567890abcdef', 10)

export default function Home() {
  const navigate = useNavigate()

  useEffect(() => {
    document.title = appName
  })

  function createBin() {
    const slug = nanoid()
    const url = `/${slug}/inspect`
    navigate(url)
  }

  return (
    <>
      <Header></Header>
      <section className='pt-20 pb-32 px-4 text-center'>
        <div className='max-w-3xl mx-auto'>
          <h1 className='text-4xl sm:text-5xl font-bold text-gray-900 mb-6'>
            Debug and Inspect HTTP Requests with Ease
          </h1>
          <p className='text-xl text-gray-600 mb-8'>
            RequestBin gives you a URL that collects requests you send it and
            lets you inspect them in a human-friendly way. Perfect for debugging
            webhooks and HTTP clients.
          </p>
          <div className='flex flex-col sm:flex-row items-center justify-center gap-4'>
            <button
              onClick={createBin}
              className='w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 cursor-pointer'
            >
              Get Started Free
              <ArrowRight className='ml-2 h-5 w-5' />
            </button>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className='bg-indigo-600 py-16'>
        <div className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center'>
          <h2 className='text-3xl font-bold text-white mb-4'>
            Ready to Debug Your HTTP Requests?
          </h2>
          <p className='text-indigo-100 mb-8 max-w-2xl mx-auto'>
            Create a free RequestBin now and start inspecting your HTTP requests
            in seconds. No signup required.
          </p>
          <button
            onClick={createBin}
            className='inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-indigo-600 bg-white hover:bg-indigo-50  cursor-pointer'
          >
            Create Your First Bin
            <ArrowRight className='ml-2 h-5 w-5' />
          </button>
        </div>
      </section>
      <Footer></Footer>
    </>
  )
}
