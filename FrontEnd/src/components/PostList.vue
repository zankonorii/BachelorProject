<script setup>
import PostItem from './PostItem.vue'
</script>

<template>
  <div class="container">
    <div v-for="post in posts" v-bind:key="post">
      <PostItem :post="post"></PostItem>
    </div>
  </div>
</template>

<script>
  import axios from "axios";

  export default{
    data(){
      return{
        posts: null
      }
    },
    async mounted(){
      await axios
          .get("http://localhost:8300/posts")
          .then(response => (this.posts = response.data))
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
