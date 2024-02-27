<template>
  <div class="home">
    <!-- <img alt="Vue logo" src="../assets/logo.png" /> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
    <v-app-bar
  color="blue-grey lighten-1"
  elevation="4"
  >
    <v-container fluid class="white">
      <v-row class="d-flex flex-row pa-3 justify-end">
          <v-btn @click="medlogin" class="ml-3">login</v-btn>        
      </v-row>
    </v-container>
    
</v-app-bar>

<div :v-show="show" >
            <v-sheet width="" height="90vh" class="d-flex" rounded="false">
                <v-row justify="center" align='center'>
                    <v-card width="400" height="300" color="">
                        <v-container fluid>
                            <v-row>
                            <v-col cols="12" md-12>
                                <v-text-field v-model="userid" label="userid">
                                </v-text-field>
                            </v-col>
                            <v-col cols="12" md-12>
                                <v-text-field v-model="password" label="password">
                                </v-text-field>
                            </v-col>
                            <v-col justify="center">
                                <v-btn md="5" @click="login" color="primary">
                                    login
                                </v-btn>
                            </v-col>
                        </v-row>
                        </v-container>
                    </v-card>
                </v-row>
            </v-sheet >
        </div>
        <!-- <div><h1>{{ role }}</h1></div> -->
        <div v-show="false" >

          <medlogin
          :role="role"
          :date="date"
          />
        </div>

  </div>

</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
// import useraddbar from "../components/useraddbar.vue"
import medlogin from "../views/medlogin.vue"

export default {
  name: "Home",
  data(){
    return{
      userid:'',
      password:'',
      show:false,
      data:this.$store.state.login,
      role:'',
      date:''
      // billeEntry:false,
      // stockEntry:false,
      // stockView:false,
      // SaleReport:false,
      // addUser:false,
      // loginHistory:false,


      // :stockView="stockView"
      //     :stockEntry="stockEntry"
      //     :SaleReport="SaleReport"
      //     :billeEntry="billeEntry"
      //     :loginHistory="loginHistory"
      //     :addUser="addUser"

    }
  },
  components: {
    // HelloWorld,
    // useraddbar,
    medlogin

  },
  methods:{
    medlogin(){
      this.show=true
    },
    login(){
      let user = this.userid
      let passw= this.password
      for (var i in this.data){
        if(this.data[i].username === user){
          if(this.data[i].password === passw){
            if(this.data[i].role==='Biller'){
              // this.stockView=true
              // this.billeEntry=true
              // this.stockEntry=false
              // this.SaleReport=false
              // this.addUser=false
              // this.loginHistory=false
              this.role='Biller'
              this.date = new Date().toLocaleString()
              this.$router.push('/medlogin')
            }
            else if(this.data[i].role==='Manager'){
              alert('user is manager')
            }
            else if(this.data[i].role === 'SystemAdmin'){
              alert('user is system admin')
            }
            else if(this.data[i].role === 'Inventry'){
              alert('user is inventry')
            }
          }
        }
      }
    }
  }
};
</script>
