package com.collegelocator.collegelocatorapplication;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.appcompat.app.AppCompatActivity;
import androidx.recyclerview.widget.RecyclerView;

import com.bumptech.glide.Glide;
import com.collegelocator.collegelocatorapplication.services.DetailsResponse;
import com.collegelocator.serviceWrapper.DetailsWrapper;
import com.xwray.groupie.GroupAdapter;
import com.xwray.groupie.Item;
import com.xwray.groupie.ViewHolder;

import interfaces.RecycleViewUpdater;

public class CollegeProfileActivity extends AppCompatActivity {

    TextView clgName, clgAbout, txtAddress, txtPhone, txtEmail, txtCutOff, txtHostel, txtAvgFees, txtClgType, txtDeemed;
    ImageView clgImage, clgLogo;

    String vrImage = null;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_college_profile);

        clgName = (TextView) findViewById(R.id.txtClgName);
        clgAbout = (TextView) findViewById(R.id.txtClgAbout);
        txtAddress = (TextView) findViewById(R.id.txtAddress);
        txtPhone = (TextView) findViewById(R.id.txtPhone);
        txtEmail = (TextView) findViewById(R.id.txtEmail);
        txtCutOff = (TextView) findViewById(R.id.txtCutOff);
        txtHostel = (TextView) findViewById(R.id.txtHostel);
        txtAvgFees = (TextView) findViewById(R.id.txtAvgFees);
        txtClgType = (TextView) findViewById(R.id.txtClgType);
        txtDeemed = (TextView) findViewById(R.id.txtDeemed);
        clgImage = (ImageView) findViewById(R.id.imgClgImage);
        clgLogo = (ImageView) findViewById(R.id.imgClgLogo);

    }

    @Override
    protected void onStart() {
        super.onStart();
        Intent intent = getIntent();
        String id = intent.getStringExtra("id");

        GroupAdapter<ViewHolder> courses = new GroupAdapter<ViewHolder>();
        RecyclerView recyclerCourses = findViewById(R.id.recyclerCourses);
        recyclerCourses.setAdapter(courses);

        Button btnVR = (Button) findViewById(R.id.btnVR);
        //Toolbar clgToolBar = (Toolbar) findViewById(R.id.clgToolBar);

        DetailsWrapper detailsWrapper = new DetailsWrapper();
        detailsWrapper.GetDetails(id, new RecycleViewUpdater() {
            @Override
            public void updateRecycleView(Object object) {
                DetailsResponse response = (DetailsResponse) object;

                clgName.setText(response.getName());
                clgAbout.setText(response.getDetails());
                txtAddress.setText(response.getAddress());
                txtPhone.setText(response.getContactNo());
                txtEmail.setText(response.getEmail());
                txtCutOff.setText(String.valueOf(response.getCutoff()));
                if (response.getHostel())
                    txtHostel.setText("Available");
                else
                    txtHostel.setText("Not Available");

                txtAvgFees.setText(String.valueOf(response.getFees()));

                if (response.getInstituteType())
                    txtClgType.setText("Government");
                else
                    txtClgType.setText("Private");

                if (response.getDeemed())
                    txtDeemed.setText("Autonomous");
                else
                    txtDeemed.setText("Not an Autonomous");


                vrImage = response.getVrImage(0);

                Glide.with(getApplicationContext())
                        .load(response.getImages(0))
                        .centerCrop()
                        .fitCenter()
                        .into(clgImage);

                Log.d("LoGo", response.getImages(1));

                //Picasso.get().load(response.getImages(1)).into(clgLogo);

                Glide.with(getApplicationContext())
                        .load(response.getImages(1))
                        .centerCrop()
                        .fitCenter()
                        .into(clgLogo);

                for (String courseName : response.getCoursesList())
                    courses.add(new Courses(courseName));


            }
        });

        btnVR.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                if(vrImage != null) {
                    Intent intent = new Intent(getApplicationContext(), VRActivity.class);
                    intent.putExtra("vr_image", vrImage);
                    startActivity(intent);
                }
            }
        });

    }
}

class Courses extends Item<ViewHolder> {
    String courseName;
    Courses(String courseName) {
        this.courseName = courseName;
    }

    @Override
    public int getLayout() {
        return R.layout.courses_single_row;
    }

    @Override
    public void bind(@NonNull ViewHolder viewHolder, int position) {
        TextView srNo = viewHolder.itemView.findViewById(R.id.txtSrNo);
        srNo.setText(String.valueOf(viewHolder.getAdapterPosition() + 1));

        TextView textCourseName = viewHolder.itemView.findViewById(R.id.txtCourseName);
        textCourseName.setText(courseName);

    }
}
