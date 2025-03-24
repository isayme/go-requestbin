import { Inbox } from 'lucide-react'

export default function Header() {
  return (
    <>
      <header className='bg-white shadow-sm'>
        <div className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4'>
          <div className='flex items-center justify-between'>
            <div className='flex items-center'>
              <Inbox className='h-8 w-8 text-indigo-600' />
              <h1 className='ml-2 text-2xl font-bold text-gray-900'>
                RequestBin
              </h1>
            </div>
            <button
              className='inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500'
              onClick={() => {}}
            >
              Create New Bin
            </button>
          </div>
        </div>
      </header>
    </>
  )
}
