<script setup lang="ts">
    defineProps<{
        username: string,
        message: string,
        time: string
    }>()
</script>

<template>
    <div class="chatbox">
        <div class="imageContainer">
            <img src="https://mdbcdn.b-cdn.net/img/Photos/Avatars/avatar-8.webp" alt="avatar" class="rounded">
        </div>
        <div class="userContainer">
            <p>{{username}}</p>
        </div>
        <div class="messageContainer">
            <p>{{message}}</p>
        </div>
        <div class="timeContainer">
            <p>{{time}}</p>
        </div>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getPost } from './api.js'

const route = useRoute()

const loading = ref(false)
const chatbox = ref(null)
const error = ref(null)

// watch the params of the route to fetch the data again
watch(() => route.params.id, fetchData, { immediate: true })

async function fetchData(id) {
  error.value = post.value = null
  loading.value = true
  
  try {
    const response = await fetch('http://localhost:8090/chats/user', {
        method: 'GET',
        credentials: 'include',
        headers: { 
            'Content-Type': 'application/json',
            'Cookie': authorization
         }
      })
      
      var chatbox.value = await response.json();
  } catch (err) {
    error.value = err.toString()
  } finally {
    loading.value = false
  }
}
</script>

<style>
    .rounded {
        border-radius: 50%;
    }
    .imageContainer {
        grid-area: image;
        width: max-content;
        height: max-content;
    }
    .userContainer {
        grid-area: username;
        align-content: center;
        padding-left: 5%;
    }
    .messageContainer {
        grid-area: message;
        align-content: center;
        padding-left: 5%;
    }
    .timeContainer {
        grid-area: time;
        align-content: center;
        align-self: self-end;
        padding-left: 5%;
    }
    .chatbox {
        display: grid;
        justify-content: stretch;
        grid-template-areas: 
            'image username username username time'
            'image message message message message';
        background-color: grey;
        width: max-content;
        height: max-content;
        padding: 10px;
    }
</style>