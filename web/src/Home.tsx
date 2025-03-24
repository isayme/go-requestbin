import {
  ArrowRight,
  Code,
  Globe,
  History,
  Lock,
  Webhook,
  Zap,
} from 'lucide-react'
import Footer from './components/Footer'
import Header from './components/Header'

export default function Home() {
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
            <button className='w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700'>
              Get Started Free
              <ArrowRight className='ml-2 h-5 w-5' />
            </button>
            <a
              href='#'
              className='w-full sm:w-auto inline-flex items-center justify-center px-6 py-3 border border-gray-300 text-base font-medium rounded-md shadow-sm text-gray-700 bg-white hover:bg-gray-50'
            >
              View Demo
            </a>
          </div>
        </div>
      </section>

      {/* Features Grid */}
      <section className='py-16 bg-white'>
        <div className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8'>
          <div className='grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8'>
            <div className='p-6 bg-white rounded-xl shadow-sm border border-gray-100'>
              <div className='w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4'>
                <Globe className='h-6 w-6 text-indigo-600' />
              </div>
              <h3 className='text-lg font-semibold text-gray-900 mb-2'>
                Public URL
              </h3>
              <p className='text-gray-600'>
                Get a unique URL instantly to start collecting HTTP requests
                from anywhere.
              </p>
            </div>

            <div className='p-6 bg-white rounded-xl shadow-sm border border-gray-100'>
              <div className='w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4'>
                <Code className='h-6 w-6 text-indigo-600' />
              </div>
              <h3 className='text-lg font-semibold text-gray-900 mb-2'>
                Request Inspector
              </h3>
              <p className='text-gray-600'>
                Examine headers, body, query parameters, and more in a clean
                interface.
              </p>
            </div>

            <div className='p-6 bg-white rounded-xl shadow-sm border border-gray-100'>
              <div className='w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4'>
                <History className='h-6 w-6 text-indigo-600' />
              </div>
              <h3 className='text-lg font-semibold text-gray-900 mb-2'>
                Request History
              </h3>
              <p className='text-gray-600'>
                Keep track of all requests with a detailed history and timeline
                view.
              </p>
            </div>

            <div className='p-6 bg-white rounded-xl shadow-sm border border-gray-100'>
              <div className='w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4'>
                <Webhook className='h-6 w-6 text-indigo-600' />
              </div>
              <h3 className='text-lg font-semibold text-gray-900 mb-2'>
                Webhook Testing
              </h3>
              <p className='text-gray-600'>
                Perfect for testing webhook integrations and API callbacks.
              </p>
            </div>

            <div className='p-6 bg-white rounded-xl shadow-sm border border-gray-100'>
              <div className='w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4'>
                <Zap className='h-6 w-6 text-indigo-600' />
              </div>
              <h3 className='text-lg font-semibold text-gray-900 mb-2'>
                Real-time Updates
              </h3>
              <p className='text-gray-600'>
                See requests appear instantly as they arrive, no refresh needed.
              </p>
            </div>

            <div className='p-6 bg-white rounded-xl shadow-sm border border-gray-100'>
              <div className='w-12 h-12 bg-indigo-100 rounded-lg flex items-center justify-center mb-4'>
                <Lock className='h-6 w-6 text-indigo-600' />
              </div>
              <h3 className='text-lg font-semibold text-gray-900 mb-2'>
                Secure by Default
              </h3>
              <p className='text-gray-600'>
                Your data is encrypted and automatically cleared after 48 hours.
              </p>
            </div>
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
          <button className='inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-indigo-600 bg-white hover:bg-indigo-50'>
            Create Your First Bin
            <ArrowRight className='ml-2 h-5 w-5' />
          </button>
        </div>
      </section>
      <Footer></Footer>
    </>
  )
}
