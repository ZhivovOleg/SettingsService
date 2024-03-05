import {$api} from "@/api/api.ts";
import {handleError} from "@/utils/handleError.ts";


export const settingsAPI = {
    async getSettings() {
        try {
            const response = await $api.get(`settings/`)
            return response.data
        } catch (err) {
            handleError(err)
        }
    },

    async addSettings(name: string) {
        try {
            return await $api.post(`settings/`, {
                serviceName: name,
                options:`{}`
            })
        } catch (err) {
            handleError(err)
        }
    },

    async updateServiceSettings(name: string, options: any) {
        try {
            return await $api.put(`settings/${name}`, {
                options: options
            })
        } catch (err) {
            handleError(err)
        }

    },

    async deleteServiceSettings(name: string) {
        try {
            return await $api.delete(`/settings/${name}`)
        } catch (err) {
            handleError(err)
        }
    },
}
