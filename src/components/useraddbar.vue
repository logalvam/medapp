
import { RouterLink } from 'vue-router';
<template>
    <div>
       <v-card width="1000">
    <v-expansion-panels focusable>
    <v-expansion-panel>
      <v-expansion-panel-header>Add</v-expansion-panel-header>
      <v-expansion-panel-content>
        <v-form >
    <v-container> 
      <v-row >
        <v-col
          cols="12"
          md="3"

        >
          <v-text-field
            v-model="userid"
            :counter="20"
            label="User Id"
            required
            class="mt-10"
          ></v-text-field>
        </v-col>

        <v-col
          cols="12"
          md="3"
        >
          <v-text-field
            v-model="password"
            :counter="20"
            class="mt-10"
            label="Pass word"
            required
          ></v-text-field>
        </v-col>

        <v-col
          cols="12"
          md="3"
        >
        <v-container id="dropdown-example-3">
    <v-overflow-btn
      class="mt-8"
      :items="dropdown_edit"
      label="Role"
      v-model="role"
    ></v-overflow-btn>
  </v-container>
        </v-col>
        <v-col cols="12" md="3" >
            <v-btn class="mt-11 ml-16" @click="add">Add</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels>
</v-card>
    <!-- <div><h1>{{ role }}</h1></div> -->
    </div>
</template>
<script>
export default{
    name:'',
    data(){
        return{
            dropdown_edit: [
        { text: 'Biller' },
        { text: 'Manage' },
        { text: 'System Admin' },
        { text: 'Inventry' },
      ],
      userid:'',
      password:'',
      role:''
      
        }
    },
    methods:{
        add(){
            let newuser = {'username':this.userid,'password':this.password,'role':this.role}
            this.$store.state.login.push(newuser)
            alert('userr added')
        }
    }
}
</script>