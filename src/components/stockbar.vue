<template>

<div>
    <v-card width="800">
        <v-dialog width="500"> 
            <template v-slot:activator="{ on, attrs }">
        <v-btn
          color="red lighten-2"
          dark
          v-bind="attrs"
          v-on="on"
        >
          add
        </v-btn>
      </template>
      <v-card width="455" height="300">
        <v-card-title  class="text-h5 grey lighten-2">
          Add Stock
        </v-card-title>
        <v-row>
            <v-col class="ml-12" cols="12" md="10">
                <v-text-field label="Brand Name" v-model="bname"></v-text-field>
            </v-col>
            <v-col class="ml-12" cols="12" md="10">
                <v-text-field label="Medicine" v-model="medicine"></v-text-field>
            </v-col>
            <v-col>
                <v-btn cols="12" md="12">Add</v-btn>
            </v-col>
        </v-row>
      </v-card>
        </v-dialog>
    <v-container> 
      <v-row >
        <v-col
          cols="12"
          md="3"

        >
          <v-text-field
            v-model="medicine"
            :counter="20"
            label="medicine"
            required
            class="mt-10"
          ></v-text-field>
        </v-col>

        <v-col
          cols="12"
          md="3"
        >
          <v-text-field
            v-model="bname"
            :counter="20"
            class="mt-10"
            label="Brand name"
            required
          ></v-text-field>
        </v-col>

        <v-col
          cols="12"
          md="3"
        >
          <v-text-field
            v-model="qty"
            :counter="20"
            class="mt-10"
            label="Qty"
            required
          ></v-text-field>
        </v-col>

        <v-col
          cols="12"
          md="3"
        >
          <v-text-field
            v-model="uprice"
            :counter="20"
            class="mt-10"
            label="uprice"
            required
          ></v-text-field>
        </v-col>

        <v-col cols="12" md="12" >
            <v-btn fluid class="mt-11">update</v-btn>
        </v-col>
      </v-row>
    </v-container>
  <!-- </v-form>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels> -->
</v-card>
</div>

</template>

<script>
export default{
    name:'',
    data(){
        return{
            bname:'',
            medicine:'',
            qty:0,
            uprice:0,
            dialog: false
        }
    }
}
</script>