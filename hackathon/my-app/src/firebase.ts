// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyAgu1zGD9VNkGleVWTPfgYBdGeJrVtWL8w",
  authDomain: "term3-mizuki-nishiyama.firebaseapp.com",
  projectId: "term3-mizuki-nishiyama",
  storageBucket: "term3-mizuki-nishiyama.appspot.com",
  messagingSenderId: "225723682059",
  appId: "1:225723682059:web:867d2539eb4c5441f4e60e"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const fireAuth = getAuth(app);