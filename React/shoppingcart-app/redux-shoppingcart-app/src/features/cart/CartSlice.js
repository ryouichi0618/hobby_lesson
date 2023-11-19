import { createSlice } from "@reduxjs/toolkit";
import cartItems from "../../cartItems";

const initialState = {
    cartItems: cartItems,
    amount: 4,
    total: 0,
}

const cartSlice = createSlice({
    name: 'cart',
    initialState,
    reducers: {
        clearCart: (state) => {
            return {cartItems : [], amount : 0, total : 0};
        },
        removeItem: (state, action) => {
            const id = action.payload;
            state.cartItems = state.cartItems.filter((item) => item.id !== id);
        },
        increase: (state, action) => {
            const id = action.payload;
            const cartItems = state.cartItems.find((item) => item.id === id);
            cartItems.amount = cartItems.amount + 1;
        },
        decrease: (state, action) => {
            const id = action.payload;
            const cartItems = state.cartItems.find((item) => item.id === id);
            cartItems.amount = cartItems.amount - 1;
        },
        calculateTotals: (state) => {
            let amount = 0;
            let total = 0;
            state.cartItems.forEach((item) => {
                amount += item.amount;
                total += item.price * item.amount;
            });
            state.amount = amount;
            state.total = total;
        },
    }
});

export const { clearCart, removeItem, increase, decrease, calculateTotals } = cartSlice.actions;
export default cartSlice.reducer;