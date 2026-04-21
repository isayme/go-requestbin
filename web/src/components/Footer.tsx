import { useMemo } from 'react'

export default function Footer() {
  const currentYear = useMemo(() => new Date().getFullYear(), [])

  return (
    <>
      <footer className='bg-gray-50 border-t border-gray-100'>
        <div className='py-6 text-center text-gray-500 text-sm'>
          © {currentYear} RequestBin. All rights reserved.
        </div>
      </footer>
    </>
  )
}
