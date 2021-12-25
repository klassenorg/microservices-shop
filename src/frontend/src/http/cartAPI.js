import {$cartService} from "./index";


export const fetchCart = async () => {
    const {data} = await $cartService.get('', {withCredentials: true})
    return data
}

export const postAddToCart = async (id, count) => {
    const {data} = await $cartService.post('add', {"product_id": ''+id, "count": count}, {withCredentials: true})
    return data
}

export const postRemoveFromCart = async (id, count) => {
    const {data} = await $cartService.post('remove', {"product_id": ''+id, "count": count}, {withCredentials: true})
    return data
}

export const postRemoveAllFromCart = async () => {
    const {data} = await $cartService.post('remove/all', null, {withCredentials: true})
    return data
}