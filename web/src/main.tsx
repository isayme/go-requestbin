import { createRoot } from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'

import Home from './Home'
import './index.css'
import Inspect from './Inspect'

createRoot(document.getElementById('root')!).render(
  <BrowserRouter>
    <Routes>
      <Route index element={<Home />} />
      <Route path=':slug/inspect' element={<Inspect />} />
    </Routes>
  </BrowserRouter>,
)
