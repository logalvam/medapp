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
                <v-btn cols="12" md="12" @click="add">Add</v-btn>
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
        <v-container id="dropdown-example-3">
              <v-overflow-btn
                class="mt-8"
                :items="medicinemaster"
                label="Medicine name"
                item-text="medicinename"
                v-model="medicine"
              ></v-overflow-btn>
            </v-container>
        </v-col>

        <v-col
          cols="12"
          md="3"
        >
        <v-container id="dropdown-example-3">
              <!-- <v-overflow-btn
                class="mt-8"
                :items="medicinemaster"
                item-text="brandname"
                label="Brand name"
                v-model="bname"
              ></v-overflow-btn> -->
              <v-text-field
              v-model="brandname1"
              :counter="20"
              class="mt-10"
              required
              disabled
            ></v-text-field>
            </v-container>
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
            <v-btn fluid class="mt-11" @click="update">update</v-btn>
        </v-col>
      </v-row>
      {{ medicinemaster }} {{ stock }}
    </v-container>
  <!-- </v-form>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-expansion-panels> -->
</v-card>
</div>

</template>

<script>
export default {
  name: "",
  data() {
    return {
      bname: "",
      medicine: "",
      qty: 0,
      brandname1:'',
      uprice: 0,
      newarray:[],
      dialog: false,
      medicinemaster: this.$store.state.medicinemaster,
      stock:this.$store.state.stock
    };
  },
  methods: {
    add() {
      let newmedicine = { medicinename: this.bname, brandname: this.medicine };
      this.$store.state.medicinemaster.push(newmedicine);
      // console.log("hai");
    },
    update(){
      let newStock = {medicinename:this.medicine,quantity:this.qty,amount:this.uprice}
      for (var i in this.stock){
        if (this.medicine === this.stock[i].medicinename){
          this.stock[i].quantity += Number(this.qty)
          this.stock[i].amount = Number(this.uprice)
        }
        this.newarray.push(this.stock[i].medicinename)
      }
      if (!this.newarray.includes(this.medicine)){

        this.$store.state.stock.push(newStock)
      }
      // let newviewStock= {medicinename:this.medicine,brandname:this.bname,quantity:this.qty,amount:this.uprice}
      // for ( var j  in this.newarray){
      //   if(this.medicine != this.stock[j].medicinename){
      //   }
      // }

      // }
      // let newmedicine = { medicinename: this.medicine, brandname: this.bname };
      // this.medicinemaster.push(newmedicine)
      alert('stock updated')
    }
  },
  watch:{
    medicine(){
      for (var i in this.medicinemaster){
        if (this.medicine=== this.medicinemaster[i].medicinename){
          this.brandname1 = this.medicinemaster[i].brandname
        }
      }
    }
  }
};
</script>