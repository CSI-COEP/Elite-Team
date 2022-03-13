package com.collegelocator.serviceWrapper;

import android.util.Log;

import com.collegelocator.collegelocatorapplication.services.CollegeLocatorServiceGrpc;
import com.collegelocator.collegelocatorapplication.services.Distance;
import com.collegelocator.collegelocatorapplication.services.Location;
import com.collegelocator.collegelocatorapplication.services.SearchRequest;
import com.collegelocator.collegelocatorapplication.services.SearchResponse;

import java.util.Iterator;
import java.util.concurrent.TimeUnit;

import interfaces.RecycleViewUpdater;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class SearchWrapper {
    ManagedChannel mSearchChannel;
    CollegeLocatorServiceGrpc.CollegeLocatorServiceBlockingStub blockingStub;

    Location location;
    Distance distance = null;

    String searchQuery;
    boolean searchQuery_not_null = false;

    boolean hostel = false;
    boolean hostel_not_null = false;

    String state = "";
    boolean state_not_null = false;

    boolean deemed = false;
    boolean deemed_not_null = false;

    int cutoff =  37;
    boolean cutoff_not_null = false;

    int fees = 39;
    boolean fees_not_null = false;

    boolean instituteType = false; // 1 = GOV and 0 = PRI
    boolean instituteType_not_null = false;

    String courseType = "";
    boolean courseType_not_null = false;

    public static String SERVICE_ADDRESS = "8.tcp.ngrok.io:13704";

    public SearchWrapper() {
        mSearchChannel = ManagedChannelBuilder.forTarget(SERVICE_ADDRESS).usePlaintext().build();
        blockingStub = CollegeLocatorServiceGrpc.newBlockingStub(mSearchChannel);
    }

    public void SearchColleges(RecycleViewUpdater updater) {
        SearchRequest.Builder requestBuilder = SearchRequest.newBuilder();
        if (searchQuery_not_null) {
            requestBuilder.setSearchQuery(getSearchQuery());
            requestBuilder.setSearchQueryIsNull(true);
        }
        if (hostel_not_null) {
            requestBuilder.setHostel(isHostel());
            requestBuilder.setHostelIsNull(true);
        }
        if (courseType_not_null) {
            requestBuilder.setCourseType(getCourseType());
            requestBuilder.setCourseTypeNotNull(true);
        }
        if (state_not_null) {
            requestBuilder.setState(getState());
            requestBuilder.setStateNull(true);
        }
        if (deemed_not_null) {
            requestBuilder.setDeemed(isDeemed());
            requestBuilder.setDeemedNull(true);
        }
        if (cutoff_not_null) {
            requestBuilder.setCutoff(getCutoff());
            requestBuilder.setCutoffNull(true);
        }
        if (fees_not_null) {
            requestBuilder.setFees(getFees());
            requestBuilder.setFeesNull(true);
        }
        if (instituteType_not_null) {
            requestBuilder.setInstituteType(isInstituteType());
            requestBuilder.setInstituteTypeNull(true);
        }
        if (location != null && distance != null) {
            requestBuilder.setLocationNull(true);
            requestBuilder.setLocation(location);
            requestBuilder.setDistance(distance);
        }

        SearchRequest searchRequest = requestBuilder.build();
        Log.d("SearchWrapper", "Search Request Is Ready...!!");
        Log.d("LOCATION_SET", "" + location);
        Log.d("DISTANCE_SET", "" + distance);
        try {
            Iterator<SearchResponse> searchResponseIterator = blockingStub.withDeadlineAfter(5, TimeUnit.MINUTES).search(searchRequest);
            while (searchResponseIterator.hasNext()) {
                SearchResponse response = searchResponseIterator.next();
                if (response != null)
                    updater.updateRecycleView(response);
            }
        } catch (Exception e) {
        }
    }

    public Location getLocation() {
        return location;
    }

    public SearchWrapper setLocation(Location location) {
        this.location = location;
        Log.d("LOCATION_SET_fun", "" + location);
        return this;
    }

    public Distance getDistance() {
        return distance;
    }

    public SearchWrapper setDistance(Distance distance) {
        this.distance = distance;
        return this;
    }

    public String getSearchQuery() {
        return searchQuery;
    }

    public SearchWrapper setSearchQuery(String searchQuery, boolean notNull) {
        this.searchQuery = searchQuery;
        this.searchQuery_not_null = notNull;
        return this;
    }

    public boolean isHostel() {
        return hostel;
    }

    public SearchWrapper setHostel(boolean hostel, boolean notNull) {
        this.hostel = hostel;
        this.hostel_not_null = notNull;
        return this;
    }

    public String getCourseType() {
        return courseType;
    }

    public SearchWrapper setCourseType(String courseType, boolean notNull) {
        this.courseType = courseType;
        this.courseType_not_null = notNull;
        return this;
    }

    public String getState() {
        return state;
    }

    public SearchWrapper setState(String state, boolean notNull) {
        this.state = state;
        this.state_not_null = notNull;
        return this;
    }

    public boolean isDeemed() {
        return deemed;
    }

    public SearchWrapper setDeemed(boolean deemed, boolean notNull) {
        this.deemed = deemed;
        this.deemed_not_null = notNull;
        return this;
    }

    public int getCutoff() {
        return cutoff;
    }

    public SearchWrapper setCutoff(int cutoff, boolean notNull) {
        this.cutoff = cutoff;
        this.cutoff_not_null = notNull;
        return this;
    }

    public int getFees() {
        return fees;
    }

    public SearchWrapper setFees(int fees, boolean notNull) {
        this.fees = fees;
        this.fees_not_null = notNull;
        return this;
    }

    public boolean isInstituteType() {
        return instituteType;
    }

    public SearchWrapper setInstituteType(boolean instituteType, boolean notNull) {
        this.instituteType = instituteType;
        this.instituteType_not_null = notNull;
        return this;
    }

    public static String getServiceAddress() {
        return SERVICE_ADDRESS;
    }

    public static void setServiceAddress(String serviceAddress) {
        SERVICE_ADDRESS = serviceAddress;
    }
}