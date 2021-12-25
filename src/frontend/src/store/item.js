import {makeAutoObservable} from "mobx";


export default class ItemStore {
    constructor() {
        this._products = []
        makeAutoObservable(this)
    }

    setProducts(items) {
        this._products = items
    }


    get products() {
        return this._products
    }

}
