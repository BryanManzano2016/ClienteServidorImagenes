package com.example.cargarimagenescliente

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.bumptech.glide.Glide
import kotlinx.android.synthetic.main.activity_main.*


class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        imagen.setOnClickListener{
            Glide.with(this).load("http://192.168.100.130:9000/").
                override(200, 200).into(imagen)
        }
    }
}
