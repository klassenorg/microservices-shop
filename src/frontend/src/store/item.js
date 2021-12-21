import {makeAutoObservable} from "mobx";


export default class ItemStore {
    constructor() {
        this._products = [
            {"id":100000,"name":"Lemon","description":"The lemon is a round, slightly elongated fruit, it has a strong and resistant skin, with an intense bright yellow colour when it is totaly ripe, giving off a special aroma when it is cut.","price":30,"image_path":"img/products/100000.jpeg"},
            {"id":100001,"name":"Apple","description":"Apples are the ideal fruit to eat at any time, having a positive role in the achievement of nourish balance. Their skin may be green, yellow or reddish, and the meat taste ranges from a bitter to sweet flavour.","price":15,"image_path":"img/products/100001.jpeg"},
            {"id":100002,"name":"Orange","description":"Orange is orange.","price":25,"image_path":"img/products/100002.jpeg"},
            {"id":100003,"name":"Qiwi","description":"Qiwi is green inside.","price":55,"image_path":"img/products/100003.jpeg"}
        ]
        makeAutoObservable(this)
    }

    setProducts(items) {
        this._products = items
    }


    get products() {
        return this._products
    }

}
