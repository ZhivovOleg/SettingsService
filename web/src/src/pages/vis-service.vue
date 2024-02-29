<template>
  <div class="content p-4 bg-gray-color">
    <div class="cards bg-white">
      <vis-add-service v-model="serviceName" @add-settings="addSettings"/>
      <vis-card v-for="element in elements"
                @click="() => selectedHandler(element)"
                :class="{'bg-select-color text-white': selected === element.id}"
                :name="element.id"/>
    </div>
    <div class="flex flex-col grow bg-white">
      <VueJSONEditor
          v-show="content.json"
          :onChange="onChangeHandler"
          :content="content.json"
          @set-instance="setInstance"
      />
      <div v-if="!content.json" class="empty h-full flex justify-center items-center border-2">
        Выберите сервис
      </div>
      <vis-menu
          v-show="content.json"
          :is-edit="isEdit"
          @on-cancel="onCancel"
          @delete-settings="openModal = true"
          @update-settings="updateSettings"
      />
      <base-modal v-if="openModal" @close="openModal = false">
        <svg class="mx-auto mb-4 text-gray-400 w-12 h-12 " aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
             fill="none" viewBox="0 0 20 20">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"/>
        </svg>
        <h3 class="mb-5 text-lg font-normal text-gray-500">Вы хотите удалить настройки сервиса?</h3>
        <div class="flex justify-around">
          <BaseButton color="warning" text="Удалить" @click="deleteSettings"/>
          <BaseButton class="ms-4" text="Отмена" @click="openModal = false"/>
        </div>
      </base-modal>
      <base-modal v-if="store.modal" @close="store.close()">
          <pre>
            {{ JSON.stringify(store.dataError, null, 2) }}
          </pre>
      </base-modal>
    </div>
  </div>
</template>

<script setup lang="ts">
import VisAddService from "@/components/vis-add-service.vue";
import VisMenu from "@/components/vis-menu.vue";
import VueJSONEditor from "@/components/vis-json-editor.vue";
import VisCard from "@/components/vis-card.vue";
import {onBeforeMount, Ref, ref, shallowRef} from "vue";
import {settingsAPI} from "@/api/settings-api.ts";
import {useNotification} from "@kyvg/vue3-notification";
import {JSONEditor} from "vanilla-jsoneditor";
import BaseModal from "@/components/shared/BaseModal.vue";
import BaseButton from "@/components/shared/BaseButton.vue";
import {useModalError} from "@/composables/useModal.ts";


interface ServiceListType {
  id: string
  json: any
  text: string | null
}

const openModal = ref<boolean>(false)
const serviceName = ref<string>()
const content = ref<ServiceListType>({} as ServiceListType)
const isEdit = ref<boolean>(false)
const instanceJSON = ref<JSONEditor | null>(null) as Ref<JSONEditor>
const elements = shallowRef<ServiceListType[]>([])
const selected = ref<string | null>(null) as Ref<string>

const store = useModalError()

const {notify} = useNotification()

onBeforeMount(() => {
  getData()
})

function setInstance(value: JSONEditor) {
  instanceJSON.value = value
}

async function getData() {
  const dataSevice = await settingsAPI.getSettings()
  elements.value = Object.entries(dataSevice).map(([key, value]: any) => ({
    id: key,
    json: JSON.parse(value),
    text: null
  })) || []
}

async function addSettings() {
  if (!serviceName.value) return
  await settingsAPI.addSettings(serviceName.value)
  serviceName.value = ''
  await getData()

  notify({
    title: "Добавлено",
    type: 'success'
  });
}


async function deleteSettings() {
  await settingsAPI.deleteServiceSettings(selected.value)
  content.value = {} as ServiceListType
  await getData()

  notify({
    title: "Удалено",
    type: 'success'
  });
  openModal.value = false
}


async function updateSettings() {
  const updateContent = instanceJSON.value.get() as ServiceListType

  if (updateContent.text) {
    await settingsAPI.updateServiceSettings(selected.value, updateContent.text)
  } else {
    await settingsAPI.updateServiceSettings(selected.value, JSON.stringify(updateContent.json))
  }
  await getData()
  isEdit.value = false

  notify({
    title: "Сохранено",
    type: 'success'
  });
}

function onChangeHandler() {
  const {contentErrors} = arguments[2]
  isEdit.value = !contentErrors;
}


function onCancel() {
  let baseValue = elements.value.find(el => el.id === selected.value) || {} as ServiceListType
  instanceJSON.value.select(null)
  instanceJSON.value.set({json: baseValue.json})
  isEdit.value = false
}


function selectedHandler(element: ServiceListType) {
  content.value = {...element}
  selected.value = element.id
}
</script>
<style lang="scss">

.content {
  overflow: hidden;
  flex-grow: 1;
  gap: 14px;
  display: flex;
}

.cards {
  min-width: 300px;
  overflow: scroll;
  width: 30%;
  border: 1px solid #e5e7eb
}

</style>
