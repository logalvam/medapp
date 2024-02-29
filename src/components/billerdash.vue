<template>
    <div>
        <v-card width="100vw" height="100vh">
            <v-container class="black mt-12" v-show="system_inventry">
                <v-row class="mt-12">
                    <v-col cols="12" md="12" >
                        <h1 class="white--text">Welcome</h1>
                    </v-col>
                </v-row>
            </v-container>
            <v-container class="black mt-12" v-show="sales">
                <v-row class="mt-12">
                    <v-col cols="12" md="12" >
                        <h1 class="white--text">TodaySales</h1>
                        <v-if></v-if>
                        <span class="white--text text--h1">{{ todaysales }} </span><v-icon color="white">mdi-arrow-down</v-icon>
                    </v-col>
                </v-row>
            </v-container>
            <v-container class="black mt-12" v-show="stockprices">
                <v-row class="mt-12">
                    <v-col cols="12" md="12" >
                        <h1 class="white--text">Current Inventry Value</h1>
                        <h1 class="white--text">{{ stockvalue }} </h1>
                    </v-col>
                </v-row>
            </v-container>
        </v-card>
        <div>
        </div>
    </div>
</template>
<script>
export default{
    name:'billdash',
    data(){
        return{
            system_inventry:false,
            billmaster:this.$store.state.billmaster,
            stock:this.$store.state.stock,
            // todaysales:0
            currentrole:this.val,
            sales:false,
            stockprices:false
        }
    },
    props:{
        val:String
    },
    methods:{
        
    },
    watch:{
        currentrole:{
            handler(){
                if(this.currentrole === 'SystemAdmin'){
                    this.system_inventry = true
                    this.sales = false
                    this.stockprices = false
                    }
                else if(this.currentrole === 'Inventry'){
                    this.system_inventry = true
                    this.sales = false
                    this.stockprices = false
                    }
                else if(this.currentrole === 'Manager'){
                    this.sales = true
                    this.stockprices = true
                    this.system_inventry = false
                    }
                else if(this.currentrole === 'Biller'){
                    this.sales = true
                    this.stockprices = false
                    this.system_inventry = true
                    } 
            },immediate:true
        }
    },
    computed:{
        todaysales(){
            let total =0
            for(var i in this.billmaster){
            let today = new Date().toLocaleDateString()
            if (today === this.billmaster[i].billdate){
                total += this.billmaster[i].netprice
            }
        }
        return total
        },
        stockvalue(){
            let total=0
            let price=0
            for (var i in this.stock){
                price= this.stock[i].amount * this.stock[i].quantity
                total +=price
                price =0
            }
            return total
        }
    },

}

</script>