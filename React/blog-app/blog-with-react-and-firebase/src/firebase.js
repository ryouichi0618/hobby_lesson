import { initializeApp } from "firebase/app";
// import { getAnalytics } from "firebase/analytics";
import {getAuth, GoogleAuthProvider} from "firebase/auth";
import { getFirestore } from "firebase/firestore";

const firebaseConfig = {
  apiKey: "AIzaSyB4LCY3YThNY07Sn_mU5gH84owuEJHgG3E",
  authDomain: "blog-cc049.firebaseapp.com",
  projectId: "blog-cc049",
  storageBucket: "blog-cc049.appspot.com",
  messagingSenderId: "985694669809",
  appId: "1:985694669809:web:ba227a1030c6a72e03d250",
  measurementId: "G-TKZ0M863LH"
};

const app = initializeApp(firebaseConfig);
// const analytics = getAnalytics(app);
const auth = getAuth(app);
const provider =new GoogleAuthProvider();
const db = getFirestore(app);

export {auth, provider, db};