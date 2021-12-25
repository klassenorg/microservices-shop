import {makeAutoObservable} from "mobx";
import {postAddToCart, postRemoveFromCart} from "../http/cartAPI";

export default class CartStore {
    constructor() {
        this._cart = {}
        makeAutoObservable(this)
    }

    setCart(items) {
        this._cart = items
    }

    addToCart(item, count) {
        postAddToCart(item, count)
        if (item in this._cart) {
            count = count + parseInt(this._cart[item])
        }
        this._cart[item] = count
    }

    removeFromCart(item) {
        if (item in this._cart) {
            postRemoveFromCart(item, 1)
            if (parseInt(this._cart[item]) > 1) {
                this._cart[item] = parseInt(this._cart[item]) - 1 + ''
            } else {
                delete this._cart[item]
            }
        }

    }

    count(item) {
        return this._cart[item]
    }

    countAll() {
        return Object.values(this._cart).reduce((partial_sum, a) => partial_sum + parseInt(a), 0)
    }

    get cart() {
        return this._cart
    }

    get totalPrice() {
        return this._totalPrice
    }

}
