<template>
    <div class="svelte-jsoneditor-vue" ref="editor" :content="content"></div>
</template>

<script setup lang="ts">
import {JSONEditor} from "vanilla-jsoneditor";
import {onMounted, Ref, ref, watch} from "vue";

// JSONEditor properties as of version 0.3.60
interface PropsType {
  content: object;
  mode: string;
  mainMenuBar: boolean;
  navigationBar: boolean;
  statusBar: boolean;
  readOnly: boolean;
  indentation: number;
  tabSize: number;
  escapeControlCharacters: boolean;
  escapeUnicodeCharacters: boolean;
  validator: (json: object) => void;
  onError: (error: Error) => void;
  onChange: (json: object) => void;
  onChangeMode: (mode: string) => void;
  onClassName: (className: string) => void;
  onRenderValue: (element: HTMLElement, json: object) => void;
  onRenderMenu: (element: HTMLElement, json: object) => void;
  queryLanguages: object;
  queryLanguageId: string;
  onChangeQueryLanguage: (queryLanguageId: string) => void;
  onFocus: () => void;
  onBlur: () => void;
}

function pickDefinedProps(propNames: Partial<PropsType>) {
  const props: Record<string, any> = {};
  for (const [ key, value ] of Object.entries(propNames)) {

    if (value !== undefined) {
      props[key] = value
    }
  }
  return props;
}
const props = withDefaults(defineProps<Partial<PropsType>>(), {
  mainMenuBar: true,
  statusBar: true,
  escapeUnicodeCharacters: true,
  escapeControlCharacters: true,
  navigationBar: true,
})

const emit = defineEmits<{
  (event:'set-instance', value: any):void
}>()

const editor = ref()
const editorInstance = ref<JSONEditor | null>(null) as Ref<JSONEditor>

watch(() => props.content, (content) => {
  editorInstance.value.select(null)
  editorInstance.value.set({json: content});
}, {deep: true})


onMounted(() => {
  editorInstance.value = new JSONEditor({
    target: editor.value,
    props: pickDefinedProps(props)
  })
  emit('set-instance', editorInstance.value)
})

</script>
<style scoped>
.svelte-jsoneditor-vue {
  min-height: 150px;
  display: flex;
  flex: 1;
}
</style>
