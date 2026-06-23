// Firebase configuration and initialization
import { initializeApp } from 'firebase/app';
import { getAnalytics, isSupported } from 'firebase/analytics';

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCx9DvDb4ZDXlzN-ecZ-lgeItYKat4phEE",
  authDomain: "lumo-cli.firebaseapp.com",
  projectId: "lumo-cli",
  storageBucket: "lumo-cli.firebasestorage.app",
  messagingSenderId: "342940383939",
  appId: "1:342940383939:web:af7927318d53787e790076",
  measurementId: "G-DX2VXGTYBN"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);

// Initialize Analytics and export it
export const initAnalytics = async () => {
  if (typeof window !== 'undefined') {
    try {
      const analyticsSupported = await isSupported();
      if (analyticsSupported) {
        return getAnalytics(app);
      }
    } catch (error) {
      console.error('Firebase Analytics error:', error);
    }
  }
  return null;
};

export default app;
