package com.collegelocator.collegelocatorapplication;

import android.content.Intent;
import android.graphics.Bitmap;
import android.graphics.drawable.Drawable;
import android.os.Bundle;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;

import com.bumptech.glide.Glide;
import com.bumptech.glide.request.target.CustomTarget;
import com.bumptech.glide.request.transition.Transition;
import com.google.vr.sdk.widgets.pano.VrPanoramaView;

public class VRActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_vractivity);

        VrPanoramaView vrPanoramaView = (VrPanoramaView) findViewById(R.id.vrPanoramaView);
        VrPanoramaView.Options options = new VrPanoramaView.Options();
        options.inputType = VrPanoramaView.Options.TYPE_MONO;


        Intent intent = getIntent();
        String DEMO_PANORAMA_LINK = intent.getStringExtra("vr_image");

        Glide.with(this).asBitmap().load(DEMO_PANORAMA_LINK).into(new CustomTarget<Bitmap>() {
            @Override
            public void onResourceReady(@NonNull Bitmap resource, @Nullable Transition<? super Bitmap> transition) {
                vrPanoramaView.loadImageFromBitmap(resource, options);
            }
            @Override
            public void onLoadCleared(@Nullable Drawable placeholder) { }
        });
    }
}
