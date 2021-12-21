import {makeAutoObservable} from "mobx";


export default class CartStore {
    constructor() {
        this._cart = {
            "100000": "5",
            "100001": "3",
            "100002": "9"
        }
        makeAutoObservable(this)
    }
    setCart(items) {
        this._cart = items
    }

    addToCart(item, count) {
        if (item in this._cart) {
            count = count + parseInt(this._cart[item])
        }
        this._cart[item] = count
    }

    removeFromCart(item) {
        if (item in this._cart) {
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

    get cart() {
        return this._cart
    }


}
