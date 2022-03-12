package com.collegelocator.serviceWrapper;

import android.util.Log;

import com.collegelocator.collegelocatorapplication.services.CollegeLocatorServiceGrpc;
import com.collegelocator.collegelocatorapplication.services.DetailsRequest;
import com.collegelocator.collegelocatorapplication.services.DetailsResponse;
import com.collegelocator.collegelocatorapplication.services.Location;

import java.util.concurrent.TimeUnit;

import interfaces.RecycleViewUpdater;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class DetailsWrapper {

    ManagedChannel mSearchChannel;
    CollegeLocatorServiceGrpc.CollegeLocatorServiceBlockingStub blockingStub;

    public static String SERVICE_ADDRESS = "2.tcp.ngrok.io:15508";

    public DetailsWrapper() {
        mSearchChannel = ManagedChannelBuilder.forTarget(SERVICE_ADDRESS).usePlaintext().build();
        blockingStub = CollegeLocatorServiceGrpc.newBlockingStub(mSearchChannel);
    }

    public void GetDetails(String collegeId, RecycleViewUpdater updater) {
        DetailsRequest detailsRequest = DetailsRequest.newBuilder()
                .setCollegeId(collegeId)
                .build();

        Log.d("DetailsWrapper", "Details Request Is Ready...!!");
        try {
            DetailsResponse response = blockingStub.withDeadlineAfter(5, TimeUnit.MINUTES).details(detailsRequest);
            if (response != null)
                updater.updateRecycleView(response);
        }catch (Exception e) {}
    }

    public static String getServiceAddress() {
        return SERVICE_ADDRESS;
    }

    public static void setServiceAddress(String serviceAddress) {
        SERVICE_ADDRESS = serviceAddress;
    }
}
