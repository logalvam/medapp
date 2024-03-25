<template>
    <div >
       <v-app-bar
 color="grey darken-3 lighten-1"
 elevation="4"
 >
   <v-container fluid class="white">
     <v-row class="d-flex flex-row pa-3 justify-end">
<div>
 <v-dialog
       transition="dialog-top-transition"
       width="400"
     >
       <template v-slot:activator="{ on, attrs }">
         <v-btn
           color="black white--text"
           v-bind="attrs"
           v-on="on"
           
         >Login</v-btn>
       </template>
       <!-- <template > -->
        <v-card>
        <v-card-title>
          <span class="text-h5">User Login</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col
                cols="12"
                sm="6"
                md="11"
              >
                <v-text-field
                  label="User name*"
                  v-model="userid"
                  required
                ></v-text-field>
              </v-col>
              <v-col
                cols="12"
                sm="6"
                md="11"
              >
                <v-text-field
                v-model="password" 
                  label="Password"
                  type="password"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          
       
          <v-btn
            color="black--text"
            text
            @click="login"
            :disabled="valid"
          >
            Login
          </v-btn>
        </v-card-actions>
      </v-card>
<!-- </template> -->
</v-dialog>
</div>
</v-row>

</v-container>
</v-app-bar>      

        <!-- <div> -->
            <v-card fluid class="grey darken-3 rounded-0"   width="100vw" height="95vh">
                <v-container>
                    

                </v-container>
            </v-card>
        <!-- </div> -->
   </div>
</template>
<script>


// import medlogin1 from "../components/medlogin.vue"
export default{
   name: 'loginpage',
   data() {
       return {
           userid: '',
           password: '',
           data: this.$store.state.login,
           date: '',
           userrole:'',
           logout:'notyet',
           valid:false,
           propuserid:'',
        
       };
   },
   methods: {
       logoutdate(value){
           this.logout=value
       },
       login() {
           let user = this.userid;
           let passw = this.password;
           this.date = new Date().toLocaleString();
           for (var i in this.data) {
               if (this.data[i].username === user) {
                   if (this.data[i].password === passw) {
                       if (this.data[i].role === 'Biller') {
                           this.userrolerole = 'Biller';
                           alert('user is biller');
                           let newuserentry = { userid: user, 'login': this.date, 'logout': this.logout };
                           this.$store.state.loginhistory.push(newuserentry)
                           this.$router.push('/medlogin');
                       }
                       else if (this.data[i].role === 'Manager') {
                               this.userrolerole = 'Manager';
                               alert('user is manager');
                                let newuserentry = { userid: user, 'login': this.date, 'logout': this.logout };
                                this.$store.state.loginhistory.push(newuserentry)
                                this.$router.push('/medlogin');
                       }
                       else if (this.data[i].role === 'SystemAdmin') {
                        this.userrolerole = 'SyatemAdmin';
                               alert('user is admin');
                                let newuserentry = { userid: user, 'login': this.date, 'logout': this.logout };
                                this.$store.state.loginhistory.push(newuserentry)
                                this.$router.push('/medlogin');
                       }
                       else if (this.data[i].role === 'Inventry') {
                           this.userrolerole = 'Inventry';
                               alert('user is Inventry');
                                let newuserentry = { userid : user, 'login': this.date, 'logout': this.logout };
                                this.$store.state.loginhistory.push(newuserentry)
                                this.$router.push('/medlogin');
                       }
                   }
                   else{
                    alert('Please verify username and password')
                   }
               }
           }
       }
   },
   components: { 
    // medlogin1
    },
    watch:{
        userid:{
            handler(){
                if(this.userid!==''){
                    this.valid=true
                }
                for (var i in this.data){
                    if(this.userid === this.data[i].username){
                        this.userrole= this.data[i].role
                        this.propuserid = this.userid
                    }
                }
            },immediate:true

        },
        password:{
            handler(){
                if(this.password!==''){
                    this.valid=false
                }
                else{
                    this.valid=true
                }
            },immediate:true
        },
                
    },
    created(){
        // console.log('created')
    },
    mounted(){
        // console.log('mounted')
    },
    destroyed(){
        // console.log('destroyed')
    }
}
</script>