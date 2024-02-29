import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    login:[
      {username:'abcd',password:'1234',role:'Biller'},
      {username:'efgh',password:'1234',role:'Biller'},
      {username:'manager',password:'1234',role:'Manager'},
      {username:'invent',password:'5678',role:'Inventry'},
      {username:'admin',password:'9012',role:'SystemAdmin'},
    ],
    loginhistory:[  
    ],
    medicinemaster:[
      {medicinename:'Bacitracin',brandname:'xyz'},
      {medicinename:'Bactrim',brandname:'xy5z'},
      {medicinename:'Aspirin',brandname:'Bayer Aspirin'},
      {medicinename:'Levothyroxine',brandname:'Levoxyl'},
      {medicinename:'Metoprolol',brandname:'Lopressor'},
    ],
    stock:[
      {medicinename:'Metoprolol',quantity:5,amount:10},
      {medicinename:'Bacitracin',quantity:3,amount:12},
      {medicinename:'Aspirin',quantity:10,amount:20},
      {medicinename:'Levothyroxine',quantity:15,amount:4},
      {medicinename:'Bactrim',quantity:6,amount:13},

    ],
    billmaster:[{billno:1,billdate:'29/02/2024',billamount:200,billgst:'18%',netprice:240,userid:'abcd'},{billno:2,billdate:'29/02/2024',billamount:100,billgst:'18%',netprice:120,userid:'efgh'},{billno:3,billdate:'29/02/2024',billamount:50,billgst:'18%',netprice:55,userid:'abcd'},{billno:4,billdate:'29/02/2024',billamount:500,billgst:'18%',netprice:640,userid:'efgh'},{billno:5,billdate:'29/02/2024',billamount:220,billgst:'18%',netprice:20,userid:'abcd'}

    ],
    billdetails:[

    ]
  },
  mutations: {},
  actions: {},
  modules: {},
});
