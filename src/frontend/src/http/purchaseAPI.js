import {$purchaseService} from "./index";


export const postPurchase = async (email, phone, name, city, address, cardNumber, cardExpiration, cardCVV) => {
    const {data} = await $purchaseService.post('', {"email": email, "phone": phone, "full_name": name, "city": city, "address": address, "card_number": cardNumber, "cvc": cardCVV, "exp": cardExpiration}, {withCredentials: true})
    return data
}

export const getOrder = async (id) => {
    const {data} = await $purchaseService.get(id, {withCredentials: true})
    return data
}