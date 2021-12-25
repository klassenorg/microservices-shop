import axios from "axios";

export const $catalogService = axios.create({
    baseURL: process.env.REACT_APP_CATALOG_SERVICE_URL
})

export const $cartService = axios.create({
    baseURL: process.env.REACT_APP_CART_SERVICE_URL
})

export const $pricingService = axios.create({
    baseURL: process.env.REACT_APP_PRICING_SERVICE_URL
})

export const $recommendationService = axios.create({
    baseURL: process.env.REACT_APP_RECOMMENDATION_SERVICE_URL
})

export const $purchaseService = axios.create({
    baseURL: process.env.REACT_APP_PURCHASE_SERVICE_URL
})