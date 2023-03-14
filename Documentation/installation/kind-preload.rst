Preload the ``cilium`` image into each worker node in the kind cluster:

.. parsed-literal::

  docker pull quay.io/go-faster/cilium:|IMAGE_TAG|
  kind load docker-image quay.io/go-faster/cilium:|IMAGE_TAG|
