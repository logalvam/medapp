<template>
    <div>
        <v-sheet width="100vw" class="">
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
                                       
                                        filled
                                        label="Filled"
                                    ></v-autocomplete>
                                </v-col>
                                <v-col cols="12" md="4">
                                    <v-text-field label="Quantity" v-model="qty" type="number"></v-text-field>
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
        <v-container></v-container>
        <v-sheet width="100vw" class="d-flex green">
            <v-container fluid>
                <v-row>
                    <v-col cols="12" md="5">
                        <v-dialog
        transition="dialog-bottom-transition"
        max-width="600"
      >
        <template v-slot:activator="{ on, attrs }">
                        <v-btn @click="print" 
                        class="black white--text mr-5"
                        v-bind="attrs"
                        v-on="on"
                        >Preview</v-btn>
                        </template>
                        <template v-slot:default="dialog">
          <v-card>
            <v-sheet >

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
            <v-row  justify="space-around"  v-for="(value,i) in tempbill    " :key="i">
                <!-- {{ value }} -->
                    <p>{{ tempbill  [i].medicinename }}</p>
                        <p>{{ tempbill  [i].qty }}</p>
                        <p>{{ tempbill  [i].unitprice }}</p>

            </v-row>
            <v-row justify=end> 
                <v-col  cols="12" md="6" >
                    <v-card class="black white--text">
                        <h4>Total:{{ total }}</h4>
                        <h4>GST:{{ gst  }}</h4>
                        <h4>NetPrice:{{ netpay }}</h4>

                    </v-card>
                </v-col>
            </v-row>
        </v-col>
    </v-row>
</v-container>
</v-sheet>
              <v-btn
                text
                @click="dialog.value = false"
              >Print</v-btn>
          </v-card>
        </template>
    </v-dialog>

                        <v-btn class="white--text black"  @click="save">Save</v-btn>
                    </v-col>
                    <v-col cols="12" md="12" class="d-flex justify-space-around">
                        <h4>Billno:{{ billnum }}</h4>
                        <h4>Date:{{ date }}</h4>
                        <h4>Total:{{ total }}</h4>
                        <h4>GST:{{ gst }}</h4>
                        <h4>Netpay:{{ netpay }}</h4>
                        <!-- {{ billdetails }} -->
                    </v-col>
                </v-row>
                <v-data-table 
                    :headers="headers"
                    :items="tempbill"
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
        billno:6,
        date:new Date().toLocaleDateString(),
        total:0,
        gst:'18%',
        netpay:0,
        bname:'',
        uprice:0,
        stock:this.$store.state.stock,
        billdetails:this.$store.state.billdetails,
      items: this.$store.state.medicinemaster,
      billmaster:this.$store.state.billmaster,
      billnum:6,
      value: null,
      tempbill:[],
      headers: [
          {
            text: 'Medicine Name',
            align: 'start',
            sortable: false,
            value: 'medicinename',
          },
          { text: 'Brand Name', value: 'brandname' },
          { text: 'Qty', value: 'qty' },
          { text: 'Unit Price', value: 'unitprice' }],
          billmasterclone:[],
          billdetailsclone:[]
    }
},
props:{
    currentuserid:String
},
methods:{
    add(){
        for(var i in this.items){
            if (this.medname === this.items[i].medicinename){
                this.bname= this.items[i].brandname
            }
        }
        // console.log(this.billdetails)
        // console.log(this.tempbill)

        // for( var j in this.stock){
        //     if (this.medname === this.stock[j].medicinename){
        //         this.uprice = Number(this.stock[j].amount)
        //         this.total +=Number(this.qty)*this.uprice
        //     }
        // }
        for(var k in this.stock){
            if(this.medname === this.stock[k].medicinename){
                if(this.qty<= this.stock[k].quantity){
                    this.uprice = this.stock[k].amount
                    this.total = Number(this.qty) * this.uprice
                    this.netpay = 18/100 *this.total  + this.total
                    this.tempbill.push({medicinename:this.medname,brandname:this.bname,qty:this.qty,unitprice:this.uprice})
                }
                else{
                    alert('stock is not available')
                    break
                }
            }
        }



    },
    print(){

    },
    save(){
        console.log(this.billno)
        console.log('llllll')
        console.log(this.billmaster)
        for(var a in this.tempbill){
            let bill = {
                billno:this.billno,
                medicinename:this.tempbill[a].medicinename,
                quantity:this.tempbill[a].qty,
                unitprice:this.tempbill[a].unitprice,
                amount:this.total
            }
            this.billdetails.push(bill)
        }
        // let date = new Date().toLocaleDateString()
        // for (var l in this.billmaster){
            //     if(this.billno === this.billmaster[l].billno){
                
                //     }
                // }
        let newbillmaster = {'billno':this.billno,'billdate':this.date,'billamount':this.total,'billgst':this.gst,'netprice':this.netpay,'userid':this.currentuserid}
        this.billmaster.push(newbillmaster)
        // console.log("................")
        // console.log(this.billmaster)
        // console.log("................")
        // console.log(this.billdetailsclone)
        // console.log(this.billdetails)
        // for (var m in this.billdetailsclone){
        //     this.billdetails.push(this.billdetailsclone[m])
    
        // }
        console.log(this.billdetails)
        console.log('billdetails after push')
        console.log(this.billmaster)
        // for (var n in this.billmasterclone){
        //     this.billmaster.push(this.billmasterclone[n])
        // }
        // console.log(this.billmaster)

        // this.billdetails.push(this.tempbill)
        this.billno +=1
        for(var i in this.items){
            if (this.medname === this.items[i].medicinename){
                this.bname= this.items[i].brandname
            }
        }
        for(var k in this.stock){
                if(this.medname=== this.stock[k].medicinename){
                    if(this.stock[k].quantity>=this.qty){
                        this.stock[k].quantity -=this.qty
                    }
                }
            }
        this.tempbill=[]
        this.netpay=0
        this.total=0
        console.log(this.billmaster)


        },
    
    },
    created(){
        // console.log(this.billmaster[0].billno)

    },
    computed:{
    }
    // watch:{
    //     qty(){
    //         for(var i in this.stock){
    //             if(this.medname=== this.stock[i].medicinename){
    //                 if(this.stock[i].quantity>this.qty){
    //                     this.stock[i].quantity -=this.qty
    //                 }
    //                 else if(this.stock[i].quantity<this.qty){
    //                     alert('stock is not available')
    //                 }
    //             }
    //         }
    //     }
    // }   
  }
</script>