<template>
     <div >
        <v-app-bar
  color="blue-grey lighten-1"
  elevation="4"
  >
    <v-container fluid class="white">
      <v-row class="d-flex flex-row pa-3 justify-end">
<div>
  <v-dialog
        transition="dialog-top-transition"
        scrollable
        width="500"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            color="primary"
            v-bind="attrs"
            v-on="on"
          >Login</v-btn>
        </template>
        <template >
          <v-card>
            <v-sheet class="d-flex" rounded="false">
                <v-row justify="center" align='center'>
                    <v-card width="" height="" color="">
                        <v-container fluid class="pt-10">
                            <v-row>
                            <v-col cols="12" md="10"  class="ml-6"> 
                                <v-text-field v-model="userid"  label="userid">
                                </v-text-field>
                            </v-col>
                            <v-col cols="12" md="10" class="ml-6" >
                                <v-text-field v-model="password" label="password">
                                </v-text-field>
                            </v-col>
                            <v-col cols="12" md="12" class="d-flex justify-center">
                                <v-btn
                                @click="login" class="" :disabled="valid" color="primary">
                                    login
                                </v-btn>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-card>
            </v-row>
        </v-sheet >
    </v-card>
</template>
</v-dialog>

</div>
</v-row>

</v-container>
</v-app-bar>      


    </div>
</template>
<script>


export default{
    name: 'loginpage',
    data() {
        return {
            userid: '',
            password: '',
            data: this.$store.state.login,
            date: '',
            logout:'notyet'
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
                            // this.role = 'Biller';
                            alert('user is biller');
                            let newuserentry = { 'userid': user, 'login': this.date, 'logout': this.logout };
                            this.$store.state.loginhistory.push(newuserentry);
                            this.$router.push('/medlogin');
                        }
                        else if (this.data[i].role === 'Manager') {
                            alert('user is manager');
                        }
                        else if (this.data[i].role === 'SystemAdmin') {
                            alert('user is system admin');
                        }
                        else if (this.data[i].role === 'Inventry') {
                            alert('user is inventry');
                        }
                    }
                }
            }
        }
    },
    components: { 

     }
}
</script>