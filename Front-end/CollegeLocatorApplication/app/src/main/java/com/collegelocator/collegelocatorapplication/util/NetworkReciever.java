package com.collegelocator.collegelocatorapplication.util;

import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.util.Log;
import android.widget.Toast;


public class NetworkReciever extends BroadcastReceiver {
    public static String status = null;

    @Override
    public void onReceive(Context context, Intent intent) {
        status = NetworkState.getConnectivityStatusString(context);
        Log.d("NetWork", status.toString());
        Toast.makeText(context, status, Toast.LENGTH_LONG).show();
        if (status == "Wifi enabled") {
            //Toast.makeText(context, "WIFI", Toast.LENGTH_LONG).show();
        } else if (status == "Mobile data enabled") {
            //Toast.makeText(context, "MOBILE", Toast.LENGTH_LONG).show();
        } else if (status == "Not connected to Internet") {
            //code when no network connected
        }
    }
}
