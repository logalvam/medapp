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
                        <span class="white--text text_h1">{{ comparesales    }} </span>
                    </v-col>
                </v-row>
            </v-container>
            <v-container class="black mt-12" v-show="manager">
                <v-row class="mt-12">
                    <v-col cols="12" md="12" >
                        <h1 class="white--text">TodaySales</h1>
                        <div v-if="yesterdaysales<comparesales">
                            <span class="red--text text_h1">{{ comparesales }} </span><v-icon color="red">mdi-arrow-down</v-icon>
                        </div>

                        <div v-if="yesterdaysales>comparesales">
                            <span class="green--text text_h1">{{ comparesales }} </span><v-icon color="green">mdi-arrow-up</v-icon>
                        </div>
                        <div v-if="comparesales===0">
                            <span class="red--text text_h1">not yet started</span><v-icon color="white">mdi-arrow-down</v-icon>
                        </div>
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
// import apexchart from 'vue-apexcharts'
export default{
    name:'billerdashboard',
    data(){
        return{
            system_inventry:false,
            billmaster:this.$store.state.billmaster,
            stock:this.$store.state.stock,
            currentrole:this.val,
            currentuser:this.user,
            sales:false,
            stockprices:false,
            manager:false,
            useramount:0,
            
        }
    },
    props:{
        val:String,
        user:String
    },
    components:{
        // apexchart
    },
    methods:{
        
    },
    watch:{
        currentrole:{
            handler(){
                if(this.currentrole === 'SystemAdmin'){
                    this.system_inventry = true
                    this.sales = false
                    this.manager = false
                    this.stockprices = false
                    }
                else if(this.currentrole === 'Inventry'){
                    this.system_inventry = true
                    this.sales = false
                    this.stockprices = false
                    this.manager = false

                    }
                else if(this.currentrole === 'Manager'){
                    this.manager = true
                    this.stockprices = true
                    this.system_inventry = false
                    this.sales = false

                    }
                else if(this.currentrole === 'Biller'){
                    this.sales = true
                    this.stockprices = false
                    this.system_inventry = false
                    this.manager = false

                    } 
            },immediate:true
        },
        currentuser:{
            handler(){
                // console.log(this.currentuser)
                for(var i in this.billmaster){
                    if(this.currentuser === this.billmaster[i].userid){
                        
                        this.useramount += this.billmaster[i].billamount
                    }
                }
            },immediate:true
        }

    },
    computed:{
        yesterdaysales(){
            let total =0
            for(var i in this.billmaster){
            let today = new Date().toLocaleDateString()
            let yesterday = new Date(today - 24 * 60 * 60 * 1000).toDateString()
            if (yesterday === this.billmaster[i].billdate ){

                // if (today > this.billmaster[i].billdate){
                    total += this.billmaster[i].netprice
                // }   
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
        },
        comparesales(){
            let amt =0
            let today = new Date().toLocaleDateString()
            for (var i in this.billmaster){
                if (today===this.billmaster[i].billdate){
                    amt += this.billmaster[i].netprice
                }
            }
            return amt
        }
    },

}

</script>
