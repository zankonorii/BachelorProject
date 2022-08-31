<template>
  <div class="login">
    <form action="#" class="form">
      <h4 class="title">Register</h4>
      <input type="text" placeholder="name" v-model="name" class="input-box">
      <input type="text" placeholder="last name" v-model="lastName" class="input-box">
      <input type="password" name="password" v-model="password" placeholder="Password" id="Password" class="input-box">
      <button type="button" class="button" @click="login()">Registration</button>
    </form>
  </div>
</template>

<script>
  import axios from "axios";

  export default{
      data(){
          return{
              name : "",
              lastName : "",
              password : "",
              danger: false,
          }
      },
      methods:{
          login(){
              axios
                  .post("http://localhost:9001/register",{
                      name : this.name,
                      password : this.password
                  })
                  .then((response) => {
                      localStorage.setItem("user", JSON.stringify(response.data));
                      this.$router.replace("/profile")
                  })
                  .catch(error => {
                      this.danger = true;
                      console.log("ERROR :: ", error);
                  })
          },
      }
  }
</script>

<style>
  @media (min-width: 1024px) {
    .login {
      min-height: 100vh;
      display: flex;
      align-items: center;
      margin-left: 20%;
    }
    
    .form{
      width: 50%;
      height: 50%;
      border-radius: 2rem;
      background-color:#344a5d;
      border: 1rem royalblue;
    }

    .danger{
      color: crimson;
      font-size: 1em;
      font-weight:bold;
      margin-left: 10%;
    }

    .title{
      color: #41b882;
      margin-left: 40%;
      margin-bottom: 1rem;
      margin-top: 1rem;
      align-content: center;
      align-items: center;
      justify-content: center;
      font-weight: 800;
      font-size: larger;
    }

    .input-box{
      width: 90%;
      border-radius: 2rem;
      border: 0 solid;
      height: 2rem;
      margin-left: 5%;
      margin-right: 5%;
      margin-top: 2rem;
      padding: 1rem;
    }

    .button{
      background-color: #41b882;
      border-radius: 50%;
      margin-left: 35%;
      margin-top: 2rem;
      width: 6rem;
      height: 2rem;
    }

  }

  
</style>
  