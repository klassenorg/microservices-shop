import {$recommendationService} from "./index";


export const fetchRecommendations = async (count) => {
    const {data} = await $recommendationService.get(count)
    return data
}