package com.collegelocator.serviceWrapper;

import android.util.Log;

import com.collegelocator.collegelocatorapplication.services.CollegeLocatorServiceGrpc;
import com.collegelocator.collegelocatorapplication.services.InitRequest;
import com.collegelocator.collegelocatorapplication.services.InitResponse;
import com.collegelocator.collegelocatorapplication.services.Location;

import java.util.Iterator;
import java.util.concurrent.TimeUnit;

import interfaces.RecycleViewUpdater;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class InitWrapper {
    ManagedChannel mSearchChannel;
    CollegeLocatorServiceGrpc.CollegeLocatorServiceBlockingStub blockingStub;

    public static String SERVICE_ADDRESS = "2.tcp.ngrok.io:15508";

    public InitWrapper() {
        mSearchChannel = ManagedChannelBuilder.forTarget(SERVICE_ADDRESS).usePlaintext().build();
        blockingStub = CollegeLocatorServiceGrpc.newBlockingStub(mSearchChannel);
    }

    public InitResponse Init() {
        InitRequest initRequest = InitRequest.newBuilder().build();
        Log.d("InitWrapper", "Init Request Is Ready...!!");
        try {
            return blockingStub.withDeadlineAfter(5, TimeUnit.MINUTES).initApp(initRequest);
        }catch (Exception e) {}
        return null;
    }

    public static String getServiceAddress() {
        return SERVICE_ADDRESS;
    }

    public static void setServiceAddress(String serviceAddress) {
        SERVICE_ADDRESS = serviceAddress;
    }

}
