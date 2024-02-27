<template>
    <div>
        <v-sheet width="600">
            <v-expansion-panels>
                <v-expansion-panel>
                    <v-expansion-panel-header>
                        Items
                    </v-expansion-panel-header>
                    <v-expansion-panel-content>
                        <v-card>
                            <v-row>
                                <v-col cols="12" md="4">
                                    <v-autocomplete
                                        v-model="medname"
                                        :items="items"
                                        item-text="medicinename"
                                        dense
                                        :filter="customFilter"
                                        filled
                                        label="Filled"
                                    ></v-autocomplete>
                                </v-col>
                                <v-col cols="12" md="4">
                                    <v-text-field label="Quantity" v-model="qty"></v-text-field>
                                </v-col>
                                <v-col cols="12" md="4">
                                    <v-btn @click="add()">Add</v-btn>
                                </v-col>
                            </v-row>
                        </v-card>
                    </v-expansion-panel-content>
                </v-expansion-panel>
            </v-expansion-panels>
        </v-sheet>
        <v-sheet width="600">
            <v-container fluid>
                <v-row>
                    <v-col cols="12" md="4">
                        <v-dialog
        transition="dialog-bottom-transition"
        max-width="600"
      >
        <template v-slot:activator="{ on, attrs }">
                        <v-btn @click="print" 
                        color="primary"
                        v-bind="attrs"
                        v-on="on"
                        >Preview</v-btn>
                        </template>
                        <template v-slot:default="dialog">
          <v-card>
            <v-sheet width="">

<v-container class="green" >
    <v-row>
        <v-col cols="12" md="12">
            <h1 class="text-center">Medical Shop Name</h1>
        </v-col>
        <v-divider></v-divider>
        <v-col cols="12" md="12" class="white">
            <v-row justify="space-around">

                <p>Medicine name</p>
                <p>Qty</p>
                <p>Amount</p>
            </v-row>
            <v-row  justify="space-around"  v-for="(value,i) in billdetails" :key="i">
                    <p>{{ billdetails[i].medicinename }}</p>
                        <p>{{ billdetails[i].qty }}</p>
                        <p>{{ billdetails[i].unitprice }}</p>

            </v-row>
        </v-col>
    </v-row>
</v-container>
</v-sheet>
              <v-btn
                text
                @click="dialog.value = false"
              >Close</v-btn>
          </v-card>
        </template>
    </v-dialog>

                        <v-btn @click="save">Save</v-btn>
                    </v-col>
                    <v-col cols="12" md="12">
                        <h4>Billno:{{ billno }}</h4>
                        <h4>Date:{{ data }}</h4>
                        <h4>Total:{{ total }}</h4>
                        <h4>GST:{{ gst }}</h4>
                        <h4>Netpay:{{ netpay }}</h4>
                        <!-- {{ billdetails }} -->
                    </v-col>
                </v-row>
                <v-data-table 
                    :headers="headers"
                    :items="billdetails"
                    :items-per-page="5"
                    class="elevation-1 black--text"
                ></v-data-table>
            </v-container>
        </v-sheet>
        <div>        
        </div>
        <div>
        </div>
    </div>
</template>
<script>
  export default {
    data (){
    return{
        medname:'',
        qty:'',
        billno:0,
        data:new Date().toLocaleDateString(),
        total:0,
        gst:0,
        netpay:0,
        bname:'',
        uprice:0,
        stock:this.$store.state.stock,
        billdetails:this.$store.state.billmaster,
      items: this.$store.state.medicinemaster,
      value: null,
      headers: [
          {
            text: 'Medicine Name',
            align: 'start',
            sortable: false,
            value: 'medicinename',
          },
          { text: 'Brand Name', value: 'brandname' },
          { text: 'Qty', value: 'qty' },
          { text: 'Unit Price', value: 'unitprice' }]
    }
},
methods:{
    add(){
        for(var i in this.items){
            if (this.medname === this.items[i].medicinename){
                this.bname= this.items[i].brandname
            }
        }
        for( var j in this.stock){
            if (this.medname === this.stock[j].medicinename){
                this.uprice = Number(this.stock[j].amount)
                this.total +=Number(this.qty)*this.uprice
            }
        }
        let newitem = {medicinename:this.medname,brandname:this.bname,qty:Number(this.qty),unitprice:this.uprice}
        this.billdetails.push(newitem)
    },
    print(){

    }
}   
  }
</script>