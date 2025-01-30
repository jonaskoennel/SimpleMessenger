<template>
    <div>
      <input type="email" v-model="email" />
      <input type="password" v-model="password" />
  
      <button @click="submit">Submit</button>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        email: ''
      }
    },
    methods: {
      submit() {
        this.$emit('submit', this.email)
      }
    }
  }
  </script>