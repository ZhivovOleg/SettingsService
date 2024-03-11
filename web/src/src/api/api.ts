import axios from 'axios'

let BASE_URL = `http://localhost:${process.env.SettingsServicePort}/api/v1/`

export const $api = axios.create({
    baseURL: BASE_URL
})