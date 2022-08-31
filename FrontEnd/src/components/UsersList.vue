<script setup>
import UserItem from './UserItem.vue'
</script>

<template>
  <div class="container">
    <div v-for="user in users" v-bind:key="user">
      <UserItem :user="user"> </UserItem>
    </div>
  </div>
</template>

<script>
  import axios from "axios";

  export default{
    data(){
      return{
        users: null
      }
    },
    async mounted(){
      await axios
          .get("http://localhost:8200/users")
          .then(response => (this.users = response.data))
    }
  }
</script>

<style scoped>
  .container {
    display: grid;
    grid-gap: 50px;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: rgba(120, 120, 120, 1);
    font-family: 'Baloo Paaji 2', cursive;
    overflow: scroll;
  }
</style>
