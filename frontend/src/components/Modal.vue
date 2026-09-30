<script setup lang="ts">
import { onMounted, ref } from 'vue';

const props = defineProps<{
  closeBtnText: string
}>();

const emit = defineEmits(['close']);

const dialogRef = ref();
onMounted(() => {
  dialogRef.value?.showModal()
})
</script>

<template>
  <dialog id="modal" ref="dialogRef" @close="$emit('close')">
    <slot />
    <button commandfor="modal" command="close">{{ closeBtnText }}</button>
  </dialog>
</template>

<style scoped></style>
