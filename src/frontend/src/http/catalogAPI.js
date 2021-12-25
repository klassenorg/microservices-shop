import {$catalogService} from "./index";

export const fetchAllProducts = async () => {
    const {data} = await $catalogService.get('')
    return data
}

export const fetchOneProduct = async (product) => {
    const {data} = await $catalogService.get(product)
    return data
}