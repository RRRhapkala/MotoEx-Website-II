function addIconToMap(coordinates, map, markerImagePath) {
  const iconStyle = new ol.style.Style({
    image: new ol.style.Icon({
      anchor: [0.5, 1],
      src: markerImagePath,
    }),
  });

  const iconFeature = new ol.Feature({
    geometry: new ol.geom.Point(coordinates),
  });

  iconFeature.setStyle(iconStyle);

  const vectorSource = new ol.source.Vector({
    features: [iconFeature],
  });

  const vectorLayer = new ol.layer.Vector({
    source: vectorSource,
  });

  map.addLayer(vectorLayer);
}
