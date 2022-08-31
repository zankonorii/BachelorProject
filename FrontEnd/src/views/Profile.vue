<template>
    <div class="login" v-if="user != null">
      <form action="#" class="form">
        <h4 class="title">Profile</h4>
        <p class="details">Name : {{user.name}}</p>
        <button type="button" class="button" @click="login()">LogIn</button>
      </form>
    </div>
    <div v-else>
      you aren't logged in, please first login
    </div>
</template>

<script>
    export default{
        data(){
            return{
                user : null,

            }
        },
        beforeCreate(){
          this.user = JSON.parse(localStorage.getItem("user"));
          console.log(this.user);

          if (this.user == null){
            this.$router.replace("/login")
          }
          console.log(this.user.name);
          // this.name = localStorage.user.name
        },
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
        width: 5rem;
        height: 2rem;
      }

      .details{
        color: #41b882;
      }

    }

    
</style>
    