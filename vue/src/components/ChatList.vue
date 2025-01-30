<script setup lang="ts">
    import Chat from './Chat.vue';
    const username = "Jonas Könnel"
    const message = "Hallo mein Name ist Jonas!"
    const time = "Just now"
</script>

<template>
    <div class="chats">
        <Chat username="jonas" message="Hallo mein Name ist Jonas!" time="Just now"/>
        <Chat username="jonas" message="Hallo mein Name ist Jonas!" time="Just now"/>
        <Chat username="jonas" message="Hallo mein Name ist Jonas!" time="Just now"/>
    </div>
</template>