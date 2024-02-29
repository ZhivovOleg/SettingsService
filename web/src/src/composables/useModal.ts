import {defineStore} from "pinia";
import {ref} from "vue";


export const useModalError = defineStore('modalError', () => {
    const modal = ref(false)
    const dataError = ref<any>(null)
    const open = (data: any) => {
        modal.value = true
        dataError.value = data
    }
    const close = () => {
        modal.value = false
        dataError.value = null
    }
    return {
        modal,
        open,
        close,
        dataError
    }
})

