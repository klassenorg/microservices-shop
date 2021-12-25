import {$pricingService} from "./index";

export const postCalculate = async () => {
    const {data} = await $pricingService.post('calculate', null, {withCredentials: true})
    return data.total_price
}