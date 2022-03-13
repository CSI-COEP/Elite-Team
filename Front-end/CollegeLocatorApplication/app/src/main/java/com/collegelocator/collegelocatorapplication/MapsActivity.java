package com.collegelocator.collegelocatorapplication;

import android.Manifest;
import android.annotation.SuppressLint;
import android.app.AlertDialog;
import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.location.Location;
import android.net.ConnectivityManager;
import android.net.Uri;
import android.os.AsyncTask;
import android.os.Bundle;
import android.provider.Settings;
import android.util.Log;
import android.view.Gravity;
import android.view.LayoutInflater;
import android.view.View;
import android.view.inputmethod.InputMethodManager;
import android.widget.ArrayAdapter;
import android.widget.Button;
import android.widget.CheckBox;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.RadioGroup;
import android.widget.SearchView;
import android.widget.Spinner;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.core.app.ActivityCompat;
import androidx.drawerlayout.widget.DrawerLayout;
import androidx.fragment.app.FragmentActivity;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.collegelocator.collegelocatorapplication.services.Distance;
import com.collegelocator.collegelocatorapplication.services.InitResponse;
import com.collegelocator.collegelocatorapplication.services.SearchResponse;
import com.collegelocator.serviceWrapper.InitWrapper;
import com.collegelocator.serviceWrapper.SearchWrapper;
import com.google.android.gms.location.FusedLocationProviderClient;
import com.google.android.gms.location.LocationServices;
import com.google.android.gms.maps.CameraUpdateFactory;
import com.google.android.gms.maps.GoogleMap;
import com.google.android.gms.maps.OnMapReadyCallback;
import com.google.android.gms.maps.SupportMapFragment;
import com.google.android.gms.maps.model.BitmapDescriptorFactory;
import com.google.android.gms.maps.model.LatLng;
import com.google.android.gms.maps.model.MarkerOptions;
import com.google.android.gms.tasks.OnSuccessListener;
import com.google.android.material.bottomsheet.BottomSheetBehavior;
import com.google.android.material.navigation.NavigationView;
import com.google.android.material.slider.LabelFormatter;
import com.google.android.material.slider.Slider;
import com.xwray.groupie.GroupAdapter;
import com.xwray.groupie.Item;
import com.xwray.groupie.OnItemClickListener;
import com.xwray.groupie.ViewHolder;

import interfaces.CallBackInterface;
import interfaces.RecycleViewUpdater;

public class MapsActivity extends FragmentActivity implements OnMapReadyCallback {

    public static final int MY_PERMISSIONS_REQUEST_LOCATION = 99;

    private GoogleMap mMap;
    private FusedLocationProviderClient fusedLocationClient;

    private InitResponse initResponse = null;
    public static Location userLocation = null;

    public CallBackInterface callBackInterface = null;
    private static GroupAdapter<ViewHolder> collegeCardAdapter;

    private SearchView searchView;
    private static SearchWrapper searchWrapper;
    private Button applyFilters;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_maps);

        searchWrapper = new SearchWrapper();

        collegeCardAdapter = new GroupAdapter<ViewHolder>();
        RecyclerView recyclerCollegeCard = findViewById(R.id.recycler_clg);
        recyclerCollegeCard.setAdapter(collegeCardAdapter);


        // Obtain the SupportMapFragment and get notified when the map is ready to be used.
        SupportMapFragment mapFragment = (SupportMapFragment) getSupportFragmentManager().findFragmentById(R.id.map);
        assert mapFragment != null;
        mapFragment.getMapAsync(this);

        /*Recycler Adapter*/
        collegeCardAdapter.setOnItemClickListener(new OnItemClickListener() {
            @Override
            public void onItemClick(@NonNull Item item, @NonNull View view) {
                Intent intent = new Intent(MapsActivity.this, CollegeProfileActivity.class);
                intent.putExtra("id", ((CollegeCard) item).id);
                startActivity(intent);
            }
        });



        searchView = (SearchView) findViewById(R.id.etSearch);
        searchView.setOnQueryTextFocusChangeListener(new View.OnFocusChangeListener() {
            @Override
            public void onFocusChange(View view, boolean hasFocus) {
                if (hasFocus) {
                    InputMethodManager imm = (InputMethodManager) view.getContext().getSystemService(Context.INPUT_METHOD_SERVICE);
                    if (imm != null) {
                        imm.showSoftInput(view, 0);
                    }
                }
            }
        });
        searchView.setOnQueryTextListener(new SearchView.OnQueryTextListener() {
            @Override
            public boolean onQueryTextSubmit(String query) {
                return false;
            }

            @Override
            public boolean onQueryTextChange(String newText) {
                String searchString = searchView.getQuery().toString();
                if (!searchString.trim().isEmpty()) {
                    searchWrapper.setSearchQuery(searchString, true);
                } else {
                    searchWrapper.setSearchQuery(searchString, false);
                }
                new SearchAsync().execute();
                return true;
            }
        });


        if (checkInternetConnection()) {
            getUserLocation();

            /*InitWrapper*/
            InitWrapper initWrapper = new InitWrapper();
            initResponse = initWrapper.Init();

            /*Navigation View*/
            NavigationView navigationView = findViewById(R.id.navigationView);
            View headerView = LayoutInflater.from(MapsActivity.this).inflate(R.layout.navigation_layout, null);
            navigationView.addHeaderView(headerView);

            DrawerLayout drawerLayout = findViewById(R.id.drawerLayout);

            Button btn_filters_navigationView = findViewById(R.id.btn_filter_navigationView);
            btn_filters_navigationView.setOnClickListener(new View.OnClickListener() {

                @Override
                public void onClick(View view) {
                    drawerLayout.openDrawer(Gravity.LEFT);
                }

            });


            /*Slider Cutoff*/
            Slider sliderCutoff = headerView.findViewById(R.id.slider_cutoff);
            sliderCutoff.setValue(initResponse.getMinRangeCutoff());
            sliderCutoff.setValueTo(initResponse.getMaxRangeCutoff());
            sliderCutoff.setValueFrom(initResponse.getMinRangeCutoff());
            sliderCutoff.setLabelFormatter(new LabelFormatter() {
                @NonNull
                @Override
                public String getFormattedValue(float value) {
                    return (String.valueOf(value) + "%");
                }
            });

            /*Slider Fees*/
            Slider sliderFees = headerView.findViewById(R.id.slider_fees);
            sliderFees.setValue(initResponse.getMinRangeFees());
            sliderFees.setValueTo(initResponse.getMaxRangeFees());
            sliderFees.setValueFrom(initResponse.getMinRangeFees());
            sliderFees.setLabelFormatter(new LabelFormatter() {
                @NonNull
                @Override
                public String getFormattedValue(float value) {
                    return (String.valueOf((int) value));
                }
            });
            /*Slider Done*/

            RadioGroup radioGroupCategory = headerView.findViewById(R.id.radioGroup0);
            radioGroupCategory.check(R.id.all_chek);

            RadioGroup radioGroupDistance = headerView.findViewById(R.id.radioGroup1);
            radioGroupDistance.check(R.id.radioBtnDAny);

            RadioGroup radioGroupInstitute = headerView.findViewById(R.id.radioGroup2);
            radioGroupInstitute.check(R.id.radioBtnIAny);

            String[] states = {"Any", "Maharashtra", "Madhya Pradesh", "Kerala", "Gujarat", "Rajasthan"};

            //spinner
            Spinner spinner = headerView.findViewById(R.id.spinner_states);
            @SuppressLint("ResourceType") ArrayAdapter<String> dataAdapter = new ArrayAdapter<String>(getBaseContext(), android.R.layout.simple_spinner_item, states);
            dataAdapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item);
            spinner.setAdapter(dataAdapter);

            applyFilters = headerView.findViewById(R.id.btn_apply_filters);
            applyFilters.setOnClickListener(new View.OnClickListener() {

                @Override
                public void onClick(View view) {

                    switch (radioGroupCategory.getCheckedRadioButtonId()) {

                        case R.id.diploma_chek:
                            searchWrapper.setCourseType("POLYTECHNIC", true);
                            break;

                        case R.id.deg_chek:
                            searchWrapper.setCourseType("ENGINEERING", true);
                            break;

                        case R.id.med_chek:
                            searchWrapper.setCourseType("MEDICAL", true);
                            break;

                        default:
                            searchWrapper.setCourseType("", false);
                    }

                    com.collegelocator.collegelocatorapplication.services.Location location1 =
                            com.collegelocator.collegelocatorapplication.services.Location.newBuilder()
                                    .setLatitude((float) userLocation.getLatitude())
                                    .setLongitude((float) userLocation.getLongitude())
                                    .build();


                    Log.d("LOCATION_SET_MAP", "" + location1);
                    switch (radioGroupDistance.getCheckedRadioButtonId()) {

                        case R.id.radioBtnNearby:
                            searchWrapper.setDistance(Distance.NEARBY);
                            searchWrapper.setLocation(location1);
                            break;

                        case R.id.radioBtnFar:
                            searchWrapper.setDistance(Distance.MID_RANGE);
                            searchWrapper.setLocation(location1);
                            break;

                        case R.id.radioBtnReallyFar:
                            searchWrapper.setDistance(Distance.LONG_RANGE);
                            searchWrapper.setLocation(location1);
                            break;

                        default:
                            searchWrapper.setDistance(null);
                            searchWrapper.setLocation(null);
                            break;
                    }

                    switch (radioGroupInstitute.getCheckedRadioButtonId()) {
                        case R.id.radioBtnGov:
                            searchWrapper.setInstituteType(false, true);
                            break;

                        case R.id.radioBtnPri:
                            searchWrapper.setInstituteType(true, true
                            );
                            break;

                        default:
                            searchWrapper.setInstituteType(false, false);
                    }

                    if (((CheckBox) headerView.findViewById(R.id.chk_deemed)).isChecked()) {
                        searchWrapper.setDeemed(true, true);
                    } else {
                        searchWrapper.setDeemed(false, false);
                    }

                    if (((CheckBox) headerView.findViewById(R.id.chk_hostel)).isChecked()) {
                        searchWrapper.setHostel(true, true);
                    } else {
                        searchWrapper.setHostel(false, false);
                    }

                    if (!states[spinner.getSelectedItemPosition()].equals("Any")) {
                        searchWrapper.setState(states[spinner.getSelectedItemPosition()], true);
                    } else {
                        searchWrapper.setState(states[spinner.getSelectedItemPosition()], false);
                    }

                    if (sliderFees.getValueFrom() != sliderFees.getValue()) {
                        searchWrapper.setFees((int) sliderFees.getValue(), true);
                    } else {
                        searchWrapper.setFees((int) sliderFees.getValue(), false);
                    }

                    if (sliderCutoff.getValueFrom() != sliderCutoff.getValue()) {
                        searchWrapper.setCutoff((int) sliderCutoff.getValue(), true);
                    } else {
                        searchWrapper.setCutoff((int) sliderCutoff.getValue(), false);
                    }

                    new SearchAsync().execute();
                }
            });

            LinearLayout bottomSheet = findViewById(R.id.persistent_bottom_sheet);
            BottomSheetBehavior<LinearLayout> bottomSheetBehavior = BottomSheetBehavior.from(bottomSheet);
            //bottomSheetBehavior.setState(BottomSheetBehavior.STATE_EXPANDED);
            bottomSheetBehavior.setBottomSheetCallback(new BottomSheetBehavior.BottomSheetCallback() {
                @Override
                public void onStateChanged(View view, int i) {
                    // do something when state changes
                    Log.d("State", "State Changed");
                }

                @Override
                public void onSlide(View view, float v) {
                    // do something when slide happens
                    Log.d("State", "State Slided");

                }
            });

            callBackInterface = new CallBackInterface() {
                @Override
                public void startTheExecutionNow() {
                    if (userLocation != null) {
                        new SearchAsync().execute();
                    }
                    Log.d("LOCATION", "LOC 1 : " + userLocation);
                }
            };

        } else {
            AlertDialog.Builder alertDialogBuilder = new AlertDialog.Builder(MapsActivity.this);
            alertDialogBuilder.setMessage("Internet Connection is not Found");
            alertDialogBuilder.setPositiveButton("Open Settings",
                    new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface arg0, int arg1) {
                            //Toast.makeText(MapsActivity.this,"You clicked yes button",Toast.LENGTH_LONG).show();
                            Intent intent = new Intent(Settings.ACTION_DATA_ROAMING_SETTINGS);
                            startActivity(intent);
                            MapsActivity.this.finish();
                        }
                    });

            alertDialogBuilder.setNegativeButton("Exit", new DialogInterface.OnClickListener() {
                @Override
                public void onClick(DialogInterface dialog, int which) {
                    ActivityCompat.finishAffinity(MapsActivity.this);
                    finish();
                }
            });

            AlertDialog alertDialog = alertDialogBuilder.create();
            alertDialog.show();
        }
    }

    class SearchAsync extends AsyncTask<com.collegelocator.collegelocatorapplication.services.Location, Void, Void> {
        @Override
        protected Void doInBackground(com.collegelocator.collegelocatorapplication.services.Location... locations) {
            runOnUiThread(new Runnable() {
                @Override
                public void run() {
                    if (locations.length != 0) {
                        searchWrapper.setLocation(locations[0]);
                    }
                    //else {
                    // searchWrapper.setLocation(null);
                    //}
                    collegeCardAdapter.clear();
                    mMap.clear();
                    showUserLocation(userLocation.getLatitude(), userLocation.getLongitude());
                    //SearchWrapper searchWrapper = new SearchWrapper();
                    searchWrapper.SearchColleges(
                            new RecycleViewUpdater() {
                                @Override
                                public void updateRecycleView(Object object) {
                                    SearchResponse response = (SearchResponse) object;
                                    Log.d("SearchWrapperTEST", "DATA : " + response);
                                    collegeCardAdapter.add(new CollegeCard(response));
                                    showCollegeLocation(response.getName(), (double) response.getLocation().getLongitude(), (double) response.getLocation().getLatitude());
                                }
                            }
                    );
                }
            });
            return (Void)null;
        }
    }

    /**
     * Manipulates the map once available.
     * This callback is triggered when the map is ready to be used.
     * This is where we can add markers or lines, add listeners or move the camera. In this case,
     * we just add a marker near Sydney, Australia.
     * If Google Play services is not installed on the device, the user will be prompted to install
     * it inside the SupportMapFragment. This method will only be triggered once the user has
     * installed Google Play services and returned to the app.
     */

    @Override
    public void onMapReady(GoogleMap googleMap) {
        mMap = googleMap;

        // Add a marker in Sydney and move the camera
        LatLng sydney = new LatLng(-34, 151);
        mMap.addMarker(new MarkerOptions().position(sydney).title("Marker in Sydney"));
        mMap.moveCamera(CameraUpdateFactory.newLatLng(sydney));

    }

    private boolean checkInternetConnection() {
        // get Connectivity Manager object to check connection
        ConnectivityManager connect = (ConnectivityManager) getSystemService(getBaseContext().CONNECTIVITY_SERVICE);

        // Check for network connections
        if (connect.getNetworkInfo(0).getState() == android.net.NetworkInfo.State.CONNECTED ||
                connect.getNetworkInfo(0).getState() == android.net.NetworkInfo.State.CONNECTING ||
                connect.getNetworkInfo(1).getState() == android.net.NetworkInfo.State.CONNECTING ||
                connect.getNetworkInfo(1).getState() == android.net.NetworkInfo.State.CONNECTED) {

            Toast.makeText(this, " Connected ", Toast.LENGTH_LONG).show();
            return true;
        } else if (
                connect.getNetworkInfo(0).getState() == android.net.NetworkInfo.State.DISCONNECTED ||
                        connect.getNetworkInfo(1).getState() == android.net.NetworkInfo.State.DISCONNECTED) {

            Toast.makeText(this, " Not Connected ", Toast.LENGTH_LONG).show();
            return false;
        }
        return false;
    }

    private void requestLocationPermission() {
        ActivityCompat.requestPermissions(MapsActivity.this, new String[]{Manifest.permission.ACCESS_FINE_LOCATION}, MY_PERMISSIONS_REQUEST_LOCATION);
    }

    private void getUserLocation() {
        fusedLocationClient = LocationServices.getFusedLocationProviderClient(MapsActivity.this);
        if (ActivityCompat.checkSelfPermission(MapsActivity.this,
                Manifest.permission.ACCESS_FINE_LOCATION
        ) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(
                MapsActivity.this,
                Manifest.permission.ACCESS_COARSE_LOCATION
        ) != PackageManager.PERMISSION_GRANTED
        ) {
            //Log.d("Loc", "No Per");
            AlertDialog.Builder alertDialogBuilder = new AlertDialog.Builder(MapsActivity.this);
            alertDialogBuilder.setTitle("Location Permission Needed");
            alertDialogBuilder.setMessage("This app needs the Location permission, please accept to use location functionality");
            alertDialogBuilder.setCancelable(false);
            alertDialogBuilder.setPositiveButton("OK",
                    new DialogInterface.OnClickListener() {
                        @Override
                        public void onClick(DialogInterface arg0, int arg1) {
                            //Toast.makeText(MapsActivity.this,"You clicked yes button",Toast.LENGTH_LONG).show();
                            requestLocationPermission();
                        }
                    });

            alertDialogBuilder.setNegativeButton("Exit", new DialogInterface.OnClickListener() {
                @Override
                public void onClick(DialogInterface dialog, int which) {
                    ActivityCompat.finishAffinity(MapsActivity.this);
                    finish();
                }
            });

            AlertDialog alertDialog = alertDialogBuilder.create();
            alertDialog.show();
        } else {
            fusedLocationClient.getLastLocation().addOnSuccessListener(MapsActivity.this, new OnSuccessListener<Location>() {
                @Override
                public void onSuccess(Location location) {
                    if (location != null) {
                        // Logic to handle location object

                        userLocation = location;
                        Log.d("LOCATION", "LOC 2 : " + userLocation);

                        Log.d("Loc", String.valueOf(location.getLatitude()));
                        Log.d("Loc", String.valueOf(location.getLongitude()));
                        showUserLocation(location.getLatitude(), location.getLongitude());
                        callBackInterface.startTheExecutionNow();
                    }
                }
            });
        }
    }

    private void showUserLocation(Double lat, Double lon) {
        LatLng user = new LatLng(lat, lon);
        mMap.addMarker(new MarkerOptions().position(user).title("You"));
        mMap.addMarker(new MarkerOptions().position(user).title("You").icon(BitmapDescriptorFactory.defaultMarker(
                BitmapDescriptorFactory.HUE_AZURE)));
        mMap.moveCamera(CameraUpdateFactory.newLatLng(user));
    }

    private void showCollegeLocation(String collegeName, Double lat, Double lon) {
        LatLng user = new LatLng(lat, lon);
        mMap.addMarker(new MarkerOptions().position(user).title(collegeName));
        mMap.moveCamera(CameraUpdateFactory.newLatLng(user));
    }


    @Override
    public void onRequestPermissionsResult(int requestCode, String permissions[], int[] grantResults) {

        if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
            // permission was granted
            if (ActivityCompat.checkSelfPermission(
                    this,
                    Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_COARSE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
                return;
            }
            fusedLocationClient.getLastLocation().addOnSuccessListener(this, new OnSuccessListener<Location>() {
                @Override
                public void onSuccess(Location location) {
                    // Got last known location. In some rare situations this can be null.
                    if (location != null) {
                        // Logic to handle location object

                        userLocation = location;
                        Log.d("LOCATION", "LOC 3 : " + userLocation);

                        Log.d("Loc", String.valueOf(location.getLatitude()));
                        Log.d("Loc", String.valueOf(location.getLongitude()));
                        showUserLocation(location.getLatitude(), location.getLongitude());
                    }
                }
            });

        } else {
            //permission denied
            Toast.makeText(this, "permission denied", Toast.LENGTH_LONG).show();
            if (!ActivityCompat.shouldShowRequestPermissionRationale(this, Manifest.permission.ACCESS_FINE_LOCATION)) {
                Intent intent = new Intent(Settings.ACTION_APPLICATION_DETAILS_SETTINGS);
                intent.setData(Uri.fromParts("package", MapsActivity.this.getApplicationInfo().packageName, null));
                startActivity(intent);
                //Intent intent = Intent(Settings.ACTION_APPLICATION_DETAILS_SETTINGS, Uri.fromParts("package", MapsActivity.this.getApplicationInfo().packageName, null),);
            }
        }
        super.onRequestPermissionsResult(requestCode, permissions, grantResults);
    }
}

class CollegeCard extends Item<ViewHolder> {
    SearchResponse searchResponse;
    String id;

    CollegeCard(SearchResponse searchResponse) {
        this.searchResponse = searchResponse;
    }

    @Override
    public int getLayout() {
        return R.layout.college_card;
    }

    @Override
    public void bind(@NonNull ViewHolder viewHolder, int position) {

        TextView textCollegeName = viewHolder.itemView.findViewById(R.id.txtClgName);
        textCollegeName.setText(searchResponse.getName());

        TextView textCollegeAddress = viewHolder.itemView.findViewById(R.id.txtClgAddress);
        textCollegeAddress.setText(searchResponse.getName());

        ImageView imageCollegeImage = viewHolder.itemView.findViewById(R.id.imgResImage);
        Glide.with(viewHolder.getRoot().getContext())
                .load(searchResponse.getImage())
                .centerCrop()
                .fitCenter()
                .into(imageCollegeImage);

        id = searchResponse.getCollegeId();

    }
}