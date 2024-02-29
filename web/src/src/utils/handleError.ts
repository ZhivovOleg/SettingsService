import {notify} from "@kyvg/vue3-notification";
import {useModalError} from "@/composables/useModal.ts";

export function handleError(err: any) {
    const store = useModalError()
    if (err.response){
        const text = `
        <b>Name</b>: ${err.response?.data} <br>
        <b>Status</b>: ${err.response.status} <br>
        <b>Headers</b>: ${err.response.headers}
        `
        notify({
            type: 'error',
            title: 'Произошла ошибка',
            text: text
        });
        store.open(err.response)
    }else {
        notify({
            type: 'error',
            title: 'Произошла ошибка',
            text: err.message
        });
    }
    throw err;
}
