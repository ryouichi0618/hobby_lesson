import Navbar from './components/Navbar';
import Modal from './components/Modal';
import './App.css';
import CartContainer from './components/CartContainer';
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { calculateTotals } from './features/cart/CartSlice';


function App() {
  const dispatch = useDispatch();
  const { cartItems } = useSelector(state => state.cart);
  const { isOpen } = useSelector(state => state.modal)
  useEffect(() => {
    dispatch(calculateTotals())
  }, [cartItems]);

  return (
    <main>
      {isOpen && <Modal />}
      <Navbar />
      <CartContainer/>
    </main>
  );
}

export default App;
