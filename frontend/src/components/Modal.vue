<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
  isOpen: boolean
  closeBtnText: string
}>();

const emit = defineEmits(['close']);

const dialogRef = ref();

watch(() => props.isOpen, (isOpen) => {
  if (!dialogRef.value) return;

  if (isOpen) {
    dialogRef.value.showModal();
  } else {
    dialogRef.value.close();
  }
});

const handleClose = () => { emit('close') };
</script>

<template>
  <dialog ref="dialogRef" @close="handleClose">
    <slot />
    <button @click="handleClose">{{ closeBtnText }}</button>
  </dialog>
</template>

<style scoped></style>
