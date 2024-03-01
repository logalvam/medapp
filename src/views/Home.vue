<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png" /> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
  
    <loginpage/>
<!-- <appexchart/> -->
 </div>

</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
// import useraddbar from "../components/useraddbar.vue"
// import medlogin from "../views/medlogin.vue"
import loginpage from "../components/loginpage.vue"
// import appexchart  from "../components/praticesappex.vue"
export default {
  name: "Home",
  data(){
    return{
      // userid:'',
      // password:'',
      // visible:false,
      // data:this.$store.state.login,
      // role:'',
      // date:'',    

      

    }
  },
  components: {
    // HelloWorld,
    // useraddbar,
    // medlogin,
    loginpage,
    // appexchart

  },
  methods:{
    // medlogin(){
    //   this.visible= !this.visible
    // },
    // login(){
    //   let user = this.userid
    //   let passw= this.password
    //   for (var i in this.data){
    //     if(this.data[i].username === user){
    //       if(this.data[i].password === passw){
    //         if(this.data[i].role==='Biller'){
    //           // this.stockView=true
    //           // this.billeEntry=true
    //           // this.stockEntry=false
    //           // this.SaleReport=false
    //           // this.addUser=false
    //           // this.loginHistory=false
    //           this.role='Biller'
    //           alert('user is biller')
    //           this.date = new Date().toLocaleString()
    //           this.$router.push('/medlogin')
    //         }
    //         else if(this.data[i].role==='Manager'){
    //           alert('user is manager')
    //         }
    //         else if(this.data[i].role === 'SystemAdmin'){
    //           alert('user is system admin')
    //         }
    //         else if(this.data[i].role === 'Inventry'){
    //           alert('user is inventry')
    //         }
    //       }
    //     }
    //   }
    // }
  },
  watch:{
    // userid:{
    //   handler(){
    //     if (this.userid===''){
    //       this.valid=true
    //     }

    //   },immediate:true
    // },
    // password:{
    //   handler(){
    //     if (this.password===''){
    //       this.valid=true
    //     }
    //     else{
    //       this.valid=false
    //     }
    //   },immediate:true
    // }
  }
};
</script>
