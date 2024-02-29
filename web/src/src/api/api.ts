import axios from 'axios'

const BASE_URL = 'http://localhost:9999/api/v1/'

export const $api = axios.create({
    baseURL: BASE_URL
})


