import React, { Suspense } from 'react';
import { lazy } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import './App.less';
// import { Login } from './routes/Login/index';

const Home = lazy(() => import('./routes/Home/index'))
const Login = lazy(() => import('./routes/Login/index'))

function App() {
  return (
    <div>demo</div>
    // <Router>
    //   <Suspense fallback={<div>Loading...</div>}>
    //     <Routes>
    //       <Route path="/" element={<Home />} />
    //       <Route path="/login" element={<Login />} />
    //     </Routes>
    //     demo
    //   </Suspense>
    // </Router>
  );
}

export default App;
