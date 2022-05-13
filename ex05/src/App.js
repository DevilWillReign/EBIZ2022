import './App.css';
import { BrowserRouter, Route, Routes } from "react-router-dom"
import Basket from './components/Basket';
import Products from './components/Products';
import Footer from './components/Footer';
import Header from './components/Header';
import Home from './components/Home';
import Login from './components/Login';
import Register from './components/Register';
import Payments from './components/Payments';

function App() {
  return (
    <>
      <Header />
      <Routes>
            <Route path='/basket' element={<Basket />} />
            <Route path='/payment' element={<Payments />} />
            <Route path='/products' element={<Products />} />
            <Route path='/' element={<Home />} />
            <Route path='/login' element={<Login />} />
            <Route path='/register' element={<Register />} />
      </Routes>
      <Footer />
    </>
  );
}

export default App;
