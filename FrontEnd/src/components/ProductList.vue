<script setup>
import ProductItem from './ProductItem.vue'
</script>

<template>
  <div class="container">
    <div v-for="product in products">
      <ProductItem :product="product"> </ProductItem>
    </div>
  </div>
</template>

<script>
  import axios from "axios";

  export default{
    data(){
      return{
        products: null
      }
    },
    async mounted(){
      await axios
          .get("http://localhost:9000/products")
          .then(response => (this.products = response.data))
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
